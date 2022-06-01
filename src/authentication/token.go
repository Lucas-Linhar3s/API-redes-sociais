package authentication

import (
	"api/src/config"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perms["usuariosOd"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)
	return token.SignedString([]byte(config.SecretKey)) // secret
}

func Validatetoken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, returnKeyVerify)
	if err != nil {
		return err
	}

	fmt.Println(token)
	return nil
}


func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	//Bearer authorizaton

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKeyVerify(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}