package models

import "github.com/golang-jwt/jwt/v5"

type Usuarios []Usuario

type Usuario struct {
	Id         int64  `json:"id"`
	Codigo     string `json:"codigo"`
	Clave      string `json:"clave"`
	Nombre     string `json:"nombre"`
	Idtipouser int64  `json:"idtipouser"`
	Tipo       string `json:"tipo"`
	Idperfil   int64  `json:"idperfil"`
	Status     string `json:"status"`
	Direccion  string `json:"direccion"`
	Direccion2 string `json:"direccion2"`
	Ciudad     string `json:"ciudad"`
	Estado     string `json:"estado"`
	Telf       string `json:"telf"`
	Cel        string `json:"cel"`
	Correo     string `json:"correo"`
	Facebook   string `json:"facebook"`
	Whatsapp   string `json:"whatsapp"`
	Instagram  string `json:"instagram"`
	Idvendedor string `json:"idvendedor"`
}

type NuevoUsuario struct {
	Id         int64  `json:"id"`
	Codigo     string `json:"codigo"`
	Clave      string `json:"clave"`
	Nombre     string `json:"nombre"`
	Idtipouser int64  `json:"idtipouser"`
	Idperfil   int64  `json:"idperfil"`
	Status     string `json:"status"`
	Direccion  string `json:"direccion"`
	Direccion2 string `json:"direccion2"`
	Ciudad     string `json:"ciudad"`
	Estado     string `json:"estado"`
	Telf       string `json:"telf"`
	Cel        string `json:"cel"`
	Correo     string `json:"correo"`
	Facebook   string `json:"facebook"`
	Whatsapp   string `json:"whatsapp"`
	Instagram  string `json:"instagram"`
	Idvendedor string `json:"idvendedor"`
}

type Token struct {
	UserID         string `json:"codigo"`
	Name           string `json:"nombre"`
	Email          string `json:"correo"`
	StandardClaims jwt.RegisteredClaims
}

type LoginUsuario struct {
	Codigo string `json:"codigo"`
	Clave  string `json:"clave"`
}

type LoginData struct {
	Status  int    `json:"status"`
	Mensaje string `json:"mensaje"`
	Token   string `json:"token"`
	User    Usuario
}

type LoginRespuesta struct {
	Status  int    `json:"status"`
	Mensaje string `json:"mensaje"`
	User    Usuario
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
