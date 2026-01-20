package database

import "github.com/golang-jwt/jwt/v5"

var jwtKey = []byte("Mangocodigo2020*")

var users = map[string]string{
	"user": "password",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Create the Signin handler

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
