package routers

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Routers struct {
	URI string
	Method string
	Func func(w http.ResponseWriter, r *http.Request)
	RequerAUTH bool
}

func Config(r *mux.Router) *mux.Router {
	routers := routersUsuarios
	routers = append(routers, routerLogin)

	for _, routers := range routers {
		
		if routers.RequerAUTH {
			r.HandleFunc(routers.URI, middleware.Logger(middleware.Autentication(routers.Func)),).Methods(routers.Method)
		} else {
		r.HandleFunc(routers.URI, middleware.Logger(routers.Func)).Methods(routers.Method)
	 	}
	}

	return r
}