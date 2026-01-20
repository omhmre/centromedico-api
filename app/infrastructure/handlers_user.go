package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

func (a *App) GetUsuarios() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		datos, err := a.DB.GetUsers()
		if err.Status > 200 {
			utils.CreateLog(err.Mensaje)
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Vendedores"
		}
		rp.Status = 10
		rp.Mensaje = "Vendedores listados correctamente!"
		// rp.Data = datos
		// log.Println(datos)
		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) PostUsuario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.NuevoUsuario
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.CreateLog(fmt.Sprintf("Error al decodificar el cuerpo de la solicitud: %v", err))
		}
		// pass, err1 := bcrypt.GenerateFromPassword([]byte(user.Clave), bcrypt.DefaultCost)
		// if err1 != nil {
		// 	utils.CreateLog(fmt.Sprintf("Error al encriptar la contraseña: %v", err1))
		// 	json.NewEncoder(w).Encode(err1)
		// }

		// user.Clave = string(pass)

		rp := a.DB.AddUsuario(user)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.LoginUsuario
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Peticion Invalida"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		data := a.DB.Login(user)
		resp := models.LoginRespuesta{Status: data.Status, Mensaje: data.Mensaje, User: data.User}
		w.Header().Add("auth-token", data.Token)
		json.NewEncoder(w).Encode(resp)
	}
}

func (a *App) ChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.LoginUsuario
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Peticion Invalida"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		data := a.DB.ChangePassword(user)
		// resp := models.LoginRespuesta{Status: data.Status, Mensaje: data.Mensaje, User: data.User}
		json.NewEncoder(w).Encode(data)
	}
}

func (a *App) DelUsuario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.Id
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
		}

		rp := a.DB.DelUsuario(user)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PutUsuario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.Usuario
		err := json.NewDecoder(r.Body).Decode(&user)
		// json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
		}

		rp := a.DB.UpdateUsuario(user)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}
