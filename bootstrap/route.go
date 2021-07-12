package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/pkg/route"
	"goblog/routes"
)

//SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router:=mux.NewRouter()
	//注册路由
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)
	return router
}