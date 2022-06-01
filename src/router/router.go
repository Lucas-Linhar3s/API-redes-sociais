package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routers.Config(r)
}