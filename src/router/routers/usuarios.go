package routers

import (
	"api/src/controllers"
	"net/http"
)

var routersUsuarios = []Routers {
	{	// CADASTRAR USUARIOS
		URI: "/usuarios",
		Method: http.MethodPost,
		Func: controllers.CreateUsers,
		RequerAUTH: false,
	},
	{	// BUSCAR USUARIOS
		URI: "/usuarios",
		Method: http.MethodGet,
		Func: controllers.SearchUsers,
		RequerAUTH: false,
	},
	{	// BUSCAR USUARIOS POR ID
		URI: "/usuarios/{userId}",
		Method: http.MethodGet,
		Func: controllers.SearchUser,
		RequerAUTH: false,
	},
	{	//ATUALIZAR USUARIOS POR ID
		URI: "/usuarios/{userId}",
		Method: http.MethodPut,
		Func: controllers.UptadeUsers,
		RequerAUTH: false,
	},
	{
		URI: "/usuarios/{userId}",
		Method: http.MethodDelete,
		Func: controllers.DeleteUsers,
		RequerAUTH: false,
	},
}