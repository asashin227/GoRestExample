package router

import (
	"../controller"
	"github.com/emicklei/go-restful"
)

func GetUserRoute() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/user/{user_id}").To(controller.FindUser).
		Doc("get a user").
		Param(ws.PathParameter("user_id", "identifier of the user").DataType("string")))
	restful.Add(ws)
}
