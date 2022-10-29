package helper

import (
	"final-project/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

func GenerateToken(payload *Token) (string, error) {
	jwtTtl, err := strconv.Atoi(config.GetEnvVariable("JWT_TTL_IN_MINUTES"))
	if err != nil {
		panic(err)
	}

	claims := jwt.MapClaims{
		"payload": payload,
		"issued":  time.Now().Add(time.Duration(jwtTtl) * time.Minute),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetEnvVariable("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
