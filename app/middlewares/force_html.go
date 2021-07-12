package middlewares

import "net/http"

// ForceHtml 强制标头返回HTML内容类型
func ForceHtml(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//设置标头
		w.Header().Set("Content-Type","text/html;charset=utf-8")
		next.ServeHTTP(w,r)
	})
}
