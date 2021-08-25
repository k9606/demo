package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/golang-module/carbon"
)

var mySigningKey = []byte("AllYourBase")

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	claims := MyCustomClaims{
		"b55ar",
		jwt.StandardClaims{
			ExpiresAt: carbon.Now().AddMinutes(1).Timestamp(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
