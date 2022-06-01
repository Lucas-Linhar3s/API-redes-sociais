package routers

import (
	"api/src/controllers"
	"net/http"
)

var routerLogin = Routers {
	URI: "/login",
	Method: http.MethodPost,
	Func: controllers.Login,
	RequerAUTH: false,
}