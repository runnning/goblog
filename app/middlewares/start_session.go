package middlewares

import (
	"goblog/pkg/session"
	"net/http"
)

//StartSession 开启session控制
func StartSession(next http.Handler)http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		//开启会话
		session.StartSession(w,r)
		//继续处理下去的请求
		next.ServeHTTP(w,r)
	})
}
