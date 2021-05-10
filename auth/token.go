package auth

import (
	"fmt"
	"me-english/utils/config"
	"me-english/utils/console"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserTokenStruct struct {
}

func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SECRETKEY)
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.SECRETKEY, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		console.Info(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.Header.Get("x-access-token")
	if keys != "" {
		return keys
	}

	bearerToken := r.Header.Get("Authorizarion")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (string, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing Method: %v", token.Header["alg"])
		}

		return config.SECRETKEY, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := fmt.Sprintf("%.0f", claims["id"])
		if err != nil {
			return "", err
		}
		return uid, nil
	}
	return "", nil
}
