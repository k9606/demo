package jwt

import (
	"demo/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang-module/carbon"
)

type MyCustomClaims struct {
	UID uint
	jwt.StandardClaims
}

func GenerateToken(UID uint) (string, error) {
	claims := MyCustomClaims{
		UID,
		jwt.StandardClaims{
			ExpiresAt: carbon.Now().AddHours(1).Timestamp(),
			Issuer:    config.GetString("app.name"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(config.GetString("app.key")))

	return ss, err
}

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("app.key")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
