package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init() {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
} 

func main() {
	config.Charge()
	r := router.Gerar()

	fmt.Println("Rodando API na url: http://localhost:3333")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
