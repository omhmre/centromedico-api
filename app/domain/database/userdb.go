package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

func (d *DB) GetUsers() ([]models.Usuario, models.Respuesta) {
	var rp models.Respuesta
	var usuarios []models.Usuario
	rows, err := d.db.Query(sqlGetUsuarios)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = err.Error()
		utils.CreateLog(err.Error())
		return usuarios, rp
	}
	defer rows.Close()
	// usuarios := []models.Usuario{}
	usuario := models.Usuario{}
	// u.id, u.codigo, u.nombre, u.idtipouser, t.tipo, u.idperfil, u.status,
	// u.direccion, u.direccion2, u.ciudad, u.estado, u.telf, u.cel, u.correo,
	// u.facebook, u.whatsapp, u.instagram, u.idvendedor
	for rows.Next() {
		rows.Scan(
			&usuario.Id,
			&usuario.Codigo,
			&usuario.Clave,
			&usuario.Nombre,
			&usuario.Idtipouser,
			&usuario.Tipo,
			&usuario.Idperfil,
			&usuario.Status,
			&usuario.Direccion,
			&usuario.Direccion2,
			&usuario.Ciudad,
			&usuario.Estado,
			&usuario.Telf,
			&usuario.Cel,
			&usuario.Correo,
			&usuario.Facebook,
			&usuario.Whatsapp,
			&usuario.Instagram,
			&usuario.Idvendedor,
			&usuario.Foto,
		)
		usuarios = append(usuarios, usuario)
	}
	rp.Status = 10
	rp.Mensaje = "Usuarios listados correctamente!"
	return usuarios, rp
}

func (d *DB) AddUsuario(i models.NuevoUsuario) models.Respuesta {
	var rp models.Respuesta
	originalClave := ""
	if i.Clave == "" {
		originalClave = strconv.Itoa(crearClave())
	} else {
		originalClave = i.Clave
	}

	// Hashear la contraseña para almacenamiento seguro
	hashedClaveBytes, errHash := bcrypt.GenerateFromPassword([]byte(originalClave), bcrypt.DefaultCost)
	if errHash != nil {
		rp.Status = 500
		rp.Mensaje = "Error al hashear la contraseña: " + errHash.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}
	hashedClave := string(hashedClaveBytes)

	// Insertar usuario en la base de datos con la contraseña hasheada
	resp, err := d.db.Exec(`INSERT INTO seguridad.usuarios (codigo,clave,nombre,idtipouser,idperfil,status,direccion,direccion2,ciudad,estado,telf,cel,correo,idvendedor,foto) VALUES
     ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15);`, i.Codigo, hashedClave, i.Nombre, i.Idtipouser, i.Idperfil, i.Status, i.Direccion, i.Direccion2, i.Ciudad, i.Estado, i.Telf, i.Cel, i.Correo, i.Idvendedor, i.Foto)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de Usuario. " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = "Usuario Agregado Correctamente."
		utils.CreateLog("AddUsuario: Usuario creado exitosamente: " + i.Codigo)
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) Login(u models.LoginUsuario) models.LoginData {
	var user models.Usuario
	var rp models.LoginData

	// 1. Fetch user by codigo to get the hashed password ONLY.
	// NO intentes comparar la clave en la consulta SQL.
	// sqlGetUsuarios debe seleccionar todos los campos necesarios, incluyendo 'clave'.
	// La consulta debería ser algo como: SELECT id, codigo, clave, nombre, ... FROM usuarios WHERE codigo = $1
	row := d.db.QueryRow(sqlGetUsuarios+" WHERE u.codigo = $1", u.Codigo)
	err := row.Scan(
		&user.Id,
		&user.Codigo,
		&user.Clave, // Esta es la clave hasheada de la DB
		&user.Nombre,
		&user.Idtipouser,
		&user.Tipo,
		&user.Idperfil,
		&user.Status,
		&user.Direccion,
		&user.Direccion2,
		&user.Ciudad,
		&user.Estado,
		&user.Telf,
		&user.Cel,
		&user.Correo,
		&user.Facebook,
		&user.Whatsapp,
		&user.Instagram,
		&user.Idvendedor,
		&user.Foto,
	)

	if err != nil {
		// sql.ErrNoRows significa que el usuario no fue encontrado
		if err == sql.ErrNoRows {
			rp.Status = 51 // Un código de estado diferente para "usuario no encontrado"
			rp.Mensaje = "Credenciales no válidas para ingresar."
			return rp
		}
		// Otros errores de la base de datos
		rp.Status = 50
		rp.Mensaje = "Error interno del servidor." + err.Error()
		return rp
	}

	// 2. Compare the provided password with the stored hash using bcrypt.
	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Clave), []byte(u.Clave))
	if errCompare != nil {
		// Las contraseñas no coinciden o hubo otro error en la comparación
		if errCompare == bcrypt.ErrMismatchedHashAndPassword {
			utils.CreateLog("Intento de login fallido (contraseña incorrecta) para usuario: " + u.Codigo)
		} else {
			utils.CreateLog("Error al comparar contraseñas para usuario " + u.Codigo + ": " + errCompare.Error())
		}
		rp.Status = 52
		rp.Mensaje = "Credenciales no válidas para ingresar."
		return rp
	}

	// Si llegamos aquí, las credenciales son correctas.

	// 3. Generar el JWT.
	duracion, err1 := strconv.Atoi(TIEMPO)
	if err1 != nil {
		utils.CreateLog("Error al convertir TIEMPO a entero: " + err1.Error())
		// Decidir cómo manejar este error. Podrías usar un valor por defecto o retornar un error.
		// Para este ejemplo, si no se puede convertir, se usará 0, lo que podría generar un token inválido.
		// Es mejor establecer una duración por defecto o hacer que la función generateJWT maneje 0 apropiadamente.
		duracion = 3600 // Valor por defecto si TIEMPO no es un número válido
	}

	rp.Status = 20
	rp.Mensaje = "Usuario permitido."
	tokenString, err2 := generateJWT(user.Codigo, duracion)
	if err2 != nil {
		rp.Status = 53 // Nuevo código de error para problemas de token
		rp.Mensaje = "Error al generar token: " + err2.Error()
		return rp
	}

	rp.Token = tokenString // Almacena el token en la respuesta
	rp.User = user         // Almacena los datos del usuario en la respuesta
	return rp
}

func generateJWT(strUsuario string, horas int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(horas) * time.Hour)
	claims := &Claims{
		Username: strUsuario,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtKey = []byte(SECRET_KEY)
	strToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return strToken, nil
}

func (d *DB) ChangePassword(u models.LoginUsuario) models.Respuesta {
	var rp models.Respuesta
	var nombre string
	// 1. Verificar que el usuario existe
	row := d.db.QueryRow(`SELECT u.nombre FROM seguridad.usuarios u WHERE u.codigo = $1;`, u.Codigo)
	err := row.Scan(&nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			rp.Status = 404
			rp.Mensaje = "Usuario no encontrado."
		} else {
			rp.Status = 500
			rp.Mensaje = "Error al consultar el usuario: " + err.Error()
		}
		return rp
	}

	// 2. Validar que se proporcionó una nueva clave
	if u.Clave == "" {
		rp.Status = 400
		rp.Mensaje = "La nueva contraseña es requerida."
		return rp
	}

	// 3. Hashear la nueva contraseña proporcionada
	hashedClaveBytes, errHash := bcrypt.GenerateFromPassword([]byte(u.Clave), bcrypt.DefaultCost)
	if errHash != nil {
		rp.Status = 500
		rp.Mensaje = "Error al hashear la nueva contraseña: " + errHash.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	hashedClave := string(hashedClaveBytes)

	// 4. Actualizar la contraseña en la base de datos
	resp, errUpdate := d.db.Exec(`UPDATE seguridad.usuarios SET clave = $1 WHERE codigo = $2;`, hashedClave, u.Codigo)
	if errUpdate != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar la contraseña: " + errUpdate.Error()
		utils.CreateLog(fmt.Sprintf("ChangePassword Error: %v", errUpdate))
		return rp
	}

	nreg, _ := resp.RowsAffected()
	if nreg > 0 {
		rp.Status = 200
		rp.Mensaje = "Contraseña actualizada exitosamente."
		utils.CreateLog(fmt.Sprintf("ChangePassword: Password updated manually for user %s", u.Codigo))
		return rp
	}

	rp.Status = 500
	rp.Mensaje = "No se pudo actualizar la contraseña."
	return rp
}

func crearClave() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	clave := r.Intn(1000000)
	return clave
}

func (d *DB) DelUsuario(i models.Id) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(`DELETE FROM seguridad.usuarios WHERE id = $1;`, i.Id)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo eliminar el usuario. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " usuario eliminado Correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) UpdateUsuario(u models.Usuario) models.Respuesta {
	var rp models.Respuesta

	// Ejecutar la consulta de actualización
	resp, err := d.db.Exec(`UPDATE seguridad.usuarios 
        SET codigo = $1, nombre = $2, idtipouser = $3, idperfil = $4, status = $5, 
            direccion = $6, direccion2 = $7, ciudad = $8, estado = $9, telf = $10, cel = $11, 
            correo = $12, facebook = $13, whatsapp = $14, instagram = $15, idvendedor = $16, foto = $17
        WHERE id = $18;`,
		u.Codigo, u.Nombre, u.Idtipouser, u.Idperfil, u.Status,
		u.Direccion, u.Direccion2, u.Ciudad, u.Estado, u.Telf, u.Cel,
		u.Correo, u.Facebook, u.Whatsapp, u.Instagram, u.Idvendedor, u.Foto, u.Id)

	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo actualizar la información del usuario. " + err.Error()
		return rp
	}

	// Verificar si se actualizó algún registro
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = "Error al verificar los registros actualizados. " + err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " usuario actualizado correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontró ningún registro con los datos proporcionados"
	}

	return rp
}
