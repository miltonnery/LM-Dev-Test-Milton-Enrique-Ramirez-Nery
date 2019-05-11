package jwt

import "github.com/dgrijalva/jwt-go"

//Claims for the JSON WEB TOKEN
type Claims struct {
	UserID   int `id:"id"`
	Username string `contracts:"username"`
	IP       string `contracts:"ip"`
	jwt.StandardClaims
}
