package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// STRINGCONECTDATABASE É A STRING DE CONEXAO COM O MYSQL
	StringConectDatabase = ""

	// PORTA ONDE A API ESTAR RODANDO
	Port = 0

	// CHAVE QUE VAI SER USADA PARA ASSINAR O TOKEN
	SecretKey []byte
)

func Charge() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	StringConectDatabase = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}