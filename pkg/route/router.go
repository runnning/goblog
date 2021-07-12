package route

import (
	"github.com/gorilla/mux"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"
)
var Router *mux.Router

//SetRoute 设置路由实例，以供Name2URL函数使用
func SetRoute(r *mux.Router) {
	Router = r
}

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return config.GetString("app.url") + url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}