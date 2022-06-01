package middleware

import (
	"api/src/authentication"
	"api/src/resposts"
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Autentication(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.Validatetoken(r); err != nil {
			resposts.Erro(w, http.StatusUnauthorized, err)
			return
		}
		nextFunc(w, r)
	}
}