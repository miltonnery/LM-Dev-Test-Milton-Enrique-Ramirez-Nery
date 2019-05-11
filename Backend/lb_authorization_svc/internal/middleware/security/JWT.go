package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"log"
)

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "signature"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		if err.Error() == "Token used before issued" {
			logrus.Warn("Please correct this inconsistency as soon as possible ->" + err.Error())
			token.Valid = true
		} else {
			return nil, false
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
