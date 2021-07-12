package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

//Guest 只允许未登入用户访问
func Guest(next HttpHandlerFunc)HttpHandlerFunc {
	return func(w http.ResponseWriter, r*http.Request) {
		if auth.Check(){
			flash.Warning("登入用户无法访问此页面")
			http.Redirect(w,r,"/",http.StatusFound)
			return
		}
		//继续接下去的请求
		next(w,r)
	}
}
