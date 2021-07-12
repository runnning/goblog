package view

import (
	"goblog/app/models/category"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

//D 是 map[string]interface{}的简写
type D map[string]interface{}

//Render 渲染通用视图
func Render(w io.Writer,data D,tplFiles ...string)  {
	RenderTemplate(w,"app",data,tplFiles...)
}

//RenderSimple 渲染简单视图
func RenderSimple(w io.Writer,data D,tplFiles ...string)  {
	RenderTemplate(w,"simple",data,tplFiles...)
}


//RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string)  {
	//通用模板数据
	var err error
	data["isLogin"]=auth.Check()
	data["loginUser"]=auth.User
	data["flash"]=flash.All()
	data["Categories"], err = category.All()

	//生成模板文件
	allFiles:=getTemplateFiles(tplFiles...)

	//解析所有模板文件
	tmpl,err:=template.New("").Funcs(template.FuncMap{
		"RouteName2URL":route.Name2URL,
	}).ParseFiles(allFiles...)
	logger.LogError(err)
	//渲染模板
	tmpl.ExecuteTemplate(w,name,data)
}

func getTemplateFiles(tplFiles ...string)[]string  {
	//设置模板相对路径
	viewDir:="resources/views/"

	//遍历传参文件列表slice，设置正确路径
	for i, f := range tplFiles {
		tplFiles[i]=viewDir+strings.Replace(f,".","/",-1)+".gohtml"
	}

	//所有布局模板文件slice
	layoutFiles,err:=filepath.Glob(viewDir+"layouts/*.gohtml")
	logger.LogError(err)

	//合并所有文件
	return append(layoutFiles,tplFiles...)
}