package database

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"math/rand"
	"net/smtp"
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
	resp, err := d.db.Exec(`INSERT INTO seguridad.usuarios (codigo,clave,nombre,idtipouser,idperfil,status,direccion,direccion2,ciudad,estado,telf,cel,correo,idvendedor) VALUES
     ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14);`, i.Codigo, hashedClave, i.Nombre, i.Idtipouser, i.Idperfil, i.Status, i.Direccion, i.Direccion2, i.Ciudad, i.Estado, i.Telf, i.Cel, i.Correo, i.Idvendedor)
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
		// Enviar correo con la contraseña ORIGINAL (sin hashear)
		subject := "Departamento de Seguridad - Nuevo Usuario Creado"
		emailBody := fmt.Sprintf("Hola %s,\n\nSe ha creado un nuevo usuario para usted en Admin.\n\nSus credenciales son:\nCódigo de Usuario: %s\nContraseña: %s\n\nPor favor, inicie sesión y cambie su contraseña lo antes posible.\n\nSaludos,\nEl equipo de Admin", i.Nombre, i.Codigo, originalClave) // Usar originalClave aquí
		mailToSend := models.MailSend{
			To:      i.Correo,
			Subject: subject,
			Body:    emailBody,
		}
		d.SendMail(mailToSend)
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " usuario Agregado Correctamente"
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
	var correo, nombre string

	// 1. Obtener correo y nombre del usuario
	row := d.db.QueryRow(`SELECT u.correo, u.nombre FROM seguridad.usuarios u WHERE u.codigo = $1;`, u.Codigo)
	err := row.Scan(&correo, &nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			rp.Status = 404
			rp.Mensaje = "Usuario no encontrado."
		} else {
			rp.Status = 500
			rp.Mensaje = "Error al consultar el usuario: " + err.Error()
		}
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	if correo != "" {
		// 2. Generar una nueva contraseña aleatoria
		originalClave := strconv.Itoa(crearClave())

		// 3. Hashear la nueva contraseña para almacenamiento seguro
		hashedClaveBytes, errHash := bcrypt.GenerateFromPassword([]byte(originalClave), bcrypt.DefaultCost)
		if errHash != nil {
			rp.Status = 500
			rp.Mensaje = "Error al hashear la nueva contraseña: " + errHash.Error()
			utils.CreateLog(rp.Mensaje)
			return rp
		}
		hashedClave := string(hashedClaveBytes)

		// 4. Actualizar la contraseña hasheada en la base de datos
		resp, errUpdate := d.db.Exec(`UPDATE seguridad.usuarios SET clave = $1 WHERE codigo = $2;`, hashedClave, u.Codigo)
		if errUpdate != nil {
			rp.Status = 500
			rp.Mensaje = "Error al actualizar la contraseña en la base de datos: " + errUpdate.Error()
			utils.CreateLog(rp.Mensaje)
			return rp
		}

		nreg, _ := resp.RowsAffected()
		if nreg == 1 {
			// 5. Enviar correo con la contraseña ORIGINAL usando el método centralizado
			subject := "Departamento de Seguridad - Cambio de Contraseña"
			emailBody := fmt.Sprintf("Hola %s,\n\nSe ha solicitado un cambio de contraseña para su usuario en Admin.\n\nSus nuevas credenciales son:\nCódigo de Usuario: %s\nContraseña: %s\n\nPor favor, inicie sesión y cambie su contraseña lo antes posible.\n\nSaludos,\nEl equipo de Admin", nombre, u.Codigo, originalClave)

			mailToSend := models.MailSend{
				To:      correo,
				Subject: subject,
				Body:    emailBody,
			}
			d.SendMail(mailToSend)

			rp.Status = 200
			rp.Mensaje = "Clave actualizada. Se ha enviado un correo con la nueva contraseña."
			return rp
		}

		rp.Status = 500
		rp.Mensaje = "No se pudo actualizar la contraseña."
		return rp

	} else {
		rp.Status = 404
		rp.Mensaje = "El usuario no tiene un correo electrónico registrado."
		return rp
	}
}

func crearClave() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	clave := r.Intn(1000000)
	return clave
}

func enviaCorreo(strEncabezado string, strHtml string, strCorreo string) {
	// Set up authentication information
	auth := smtp.PlainAuth("", "cio@nanocodigo.com", "mangocodigo2020*", "mail.nanocodigo.com")

	// Set up the SMTP server
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "mail.nanocodigo.com",
	}
	conn, err := tls.Dial("tcp", "mail.nanocodigo.com:465", tlsconfig)
	if err != nil {
		utils.CreateLog(err.Error())
		return
	}
	defer conn.Close()

	// Create a new SMTP client
	client, err := smtp.NewClient(conn, "mail.nanocodigo.com")
	if err != nil {
		utils.CreateLog(err.Error())
		return
	}
	defer client.Quit()

	// Authenticate the client
	if err := client.Auth(auth); err != nil {
		utils.CreateLog(err.Error())
		return
	}

	// Set up the message
	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: Sistema de Correo de NanoCodigo\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
		"%s\r\n\r\n"+
		"%s", strCorreo, strEncabezado, strHtml)

	// Send the message
	if err := client.Mail("cio@nanocodigo.com"); err != nil {
		utils.CreateLog(err.Error())
		return
	}
	if err := client.Rcpt(strCorreo); err != nil {
		utils.CreateLog(err.Error())
		return
	}
	w, err := client.Data()
	if err != nil {
		utils.CreateLog(err.Error())
		return
	}
	_, err = w.Write([]byte(msg))
	if err != nil {
		utils.CreateLog(err.Error())
		return
	}
	err = w.Close()
	if err != nil {
		utils.CreateLog(err.Error())
		return
	}
}

func enviaCorreoGmail(strEncabezado string, strHtml string, strCorreo string) {
	// Información de autenticación para Gmail
	from := "omhmre@gmail.com"        // Consider making this configurable
	password := "gtct foke zsxd abaw" // ¡Reemplaza con tu contraseña real!
	to := []string{strCorreo}

	// Servidor SMTP de Gmail
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Dirección del servidor SMTP
	addr := smtpHost + ":" + smtpPort

	// Configuración de la autenticación
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Construcción del mensaje
	msg := []byte("To: " + strCorreo + "\r\n" +
		"Subject: " + strEncabezado + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
		strHtml + "\r\n")

	// Envío del correo electrónico
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		fmt.Println("Error al enviar el correo:", err)
		// Aquí puedes agregar tu lógica de logging de errores (reemplazando utils.CreateLog)
		return
	}

	fmt.Println("Correo enviado exitosamente a:", strCorreo)
	// Aquí puedes agregar tu lógica de logging de éxito
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
            correo = $12, facebook = $13, whatsapp = $14, instagram = $15, idvendedor = $16
        WHERE id = $17;`,
		u.Codigo, u.Nombre, u.Idtipouser, u.Idperfil, u.Status,
		u.Direccion, u.Direccion2, u.Ciudad, u.Estado, u.Telf, u.Cel,
		u.Correo, u.Facebook, u.Whatsapp, u.Instagram, u.Idvendedor, u.Id)

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
