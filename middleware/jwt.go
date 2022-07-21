package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type dataUser struct {
	Id   float64
	Nama string
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return " "
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	requiredSecret := os.Getenv("SECRET_KEY_JWT")

	if requiredSecret == "" {
		log.Fatal("Silahkan set SECRET_KEY_JWT DI .env")

	}

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(requiredSecret), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil
	}
	return nil
}
func ExtractTokenMetadata(r *http.Request) (*dataUser, error) {
	token, err := VerifyToken(r)

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, ok := claims["id"].(float64)
		if !ok {
			return nil, err
		}
		nama, ok := claims["nama"].(string)
		if !ok {
			return nil, err
		}
		return &dataUser{
			Id:   id,
			Nama: nama,
		}, nil
	}
	return nil, err
}
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenAuth, err := ExtractTokenMetadata(c.Request)

		if err != nil {
			responKesalahan(c, 401, "Gagal validasi JWT Token")
			return
		}

		c.Set("id", tokenAuth.Id)
		c.Set("nama", tokenAuth.Nama)

		c.Next()
	}
}
