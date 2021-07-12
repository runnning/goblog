package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	middlewares2 "goblog/app/http/middlewares"
	"goblog/app/middlewares"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router)  {
	//静态页面
	pc:=new(controllers.PagesController)
	r.NotFoundHandler=http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/about",pc.About).Methods("GET").Name("about")

	//文章相关页面
	ac:=new(controllers.ArticlesController)
	r.HandleFunc("/",ac.Index).Methods("GET").Name("home")
	r.HandleFunc("/articles/{id:[0-9]+}",ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles",middlewares2.Auth(ac.Store)).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/create",middlewares2.Auth(ac.Create)).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles/{id:[0-9]+}/edit",middlewares2.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}",middlewares2.Auth(ac.Update)).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete",middlewares2.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	//静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
	//中间件:强制内容类型为HTML
	//r.Use(middlewares.ForceHtml)

	//用户认证
	auc:=new(controllers.AuthController)
	r.HandleFunc("/auth/register",middlewares2.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register",middlewares2.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")
	r.HandleFunc("/auth/login",middlewares2.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin",middlewares2.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout",middlewares2.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	uc:=new(controllers.UserController)
	r.HandleFunc("/users/{id:[0-9]+}",uc.Show).Methods("GET").Name("users.show")

	//文章分类
	cc:=new(controllers.CategoriesController)
	r.HandleFunc("/categories/create",middlewares2.Auth(cc.Create)).Methods("GET").Name("categories.create")
	r.HandleFunc("/categories",middlewares2.Auth(cc.Store)).Methods("POST").Name("categories.store")
	r.HandleFunc("/categories/{id:[0-9]+}",cc.Show).Methods("GET").Name("categories.show")
	//全局中间件
	//开始会话
	r.Use(middlewares.StartSession)
}