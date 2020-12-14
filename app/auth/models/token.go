package models

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID string
	Name   string
	Email  string
	*jwt.StandardClaims
}
