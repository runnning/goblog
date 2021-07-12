package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	config2 "goblog/pkg/config"
	"goblog/pkg/database"
	"net/http"
)

var router *mux.Router
var db *sql.DB

func init()  {
	//初始化配置信息
	config.Initialize()
}
func main() {
	database.Initialize()
	db=database.DB

	bootstrap.SetupDB()
	router=bootstrap.SetupRoute()

	http.ListenAndServe(":"+config2.GetString("app.port"), middlewares.RemoveTrailingSlash(router))

}