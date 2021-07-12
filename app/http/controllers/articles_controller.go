package controllers

import (
	"database/sql"
	"fmt"
	"goblog/app/models/article"
	"goblog/app/polices"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
	BaseController
}

//ArticlesFormData  创建博文表单数据
type ArticlesFormData struct {
	Title,Body string
	Article article.Article
	Errors map[string]string
}

var db *sql.DB


//Show 文章详情页
func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 读取对应的文章数据
	_article, err :=article.Get(id)
	//如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w,err)
	} else {
		//读取成功，显示文章
		view.Render(w, view.D{
			"Article":_article,
			"CanModifyArticle":polices.CanModifyArticle(_article),
		}, "articles.show","articles._article_meta")
	}
}

//Index 文章列表
func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	// 1. 获取结果集
	articles,err:=article.GetAll()

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {

		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles":  articles,
		}, "articles.index", "articles._article_meta")
	}
}

//Create 文章创建页面
func (*ArticlesController) Create(w http.ResponseWriter,r *http.Request)  {
	view.Render(w,view.D{},"articles.create","articles._form_field")
}

//Store 文章创建页面
func (*ArticlesController) Store(w http.ResponseWriter,r *http.Request)  {

	//初始化数据
	currentUser:=auth.User()
	_article:=article.Article{
		Title:r.PostFormValue("title"),
		Body: r.PostFormValue("body"),
		UserID: currentUser.ID,
	}
	//表单验证
	errors:=requests.ValidateArticleForm(_article)
	//检测错误
	if len(errors)==0{
		//创建文章
		_article.Create()
		if _article.ID>0{
			indexURL:=route.Name2URL("articles.show","id",_article.GetStringID())
			http.Redirect(w,r,indexURL,http.StatusFound)
		}else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w,"创建文章失败，请联系管理员")
		}
	}else {
		view.Render(w,view.D{
			"Article":_article,
			"Errors":errors,
		},"articles.create","articles._form_field")
	}
}

//Edit 文章更新页面
func (ac *ArticlesController)Edit(w http.ResponseWriter,r *http.Request)  {
	//获取url参数
	id:=route.GetRouteVariable("id",r)
	//读取对应的文章数据
	_article,err:=article.Get(id)
	//如果出现错误
	if err!=nil{
		ac.ResponseForSQLError(w,err)
	}else {
		//检查权限
		if !polices.CanModifyArticle(_article){
			ac.ResponseForUnauthorized(w,r)
		}else {
			view.Render(w,view.D{
				"Article":_article,
				"Errors":view.D{},
			},"articles.edit","articles._form_field")
		}
	}
}

//Update 更新文章
func (ac *ArticlesController)Update(w http.ResponseWriter,r *http.Request)  {
	//获取url参数
	id:=route.GetRouteVariable("id",r)
	//读取对应的文章数据
	_article,err:=article.Get(id)
	//如果出现错误
	if err!=nil{
		ac.ResponseForSQLError(w,err)
	}else {
		//未出现错误

		//检查权限
		if !polices.CanModifyArticle(_article){
			ac.ResponseForUnauthorized(w,r)
		}else {
			//表单验证
			_article.Title=r.PostFormValue("title")
			_article.Body=r.PostFormValue("body")

			errors:=requests.ValidateArticleForm(_article)

			if len(errors)==0 {
				//表单验证通过,更新数据
				rowsAffected,err:=_article.Update()

				if err!=nil{
					//数据库错误
					logger.LogError(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w,"500服务器内部错误")
				}
				//更新成功,跳转到文章详情页
				if rowsAffected>0{
					showUrl:=route.Name2URL("articles.show","id",id)
					http.Redirect(w,r,showUrl,http.StatusFound)
				}else {
					fmt.Fprint(w,"你没有做任何更改!")
				}
			}else {
				//表单验证没通过，显示理由
				view.Render(w,view.D{
					"Article":_article,
					"Errors":errors,
				},"articles.edit","articles._form_field")
			}
		}
	}
}

//Delete 删除文章
func (ac *ArticlesController)Delete(w http.ResponseWriter,r *http.Request)  {

	//获取url参数
	id:=route.GetRouteVariable("id",r)
	_article,err:=article.Get(id)

	//如果出现错误
	if err!=nil{
		ac.ResponseForSQLError(w,err)
	}else {
		//未出现错误，执行删除操作

		//检查权限
		if !polices.CanModifyArticle(_article){
			ac.ResponseForUnauthorized(w,r)
		}else {
			//未出现错误，执行删除操作
			rowsAffected,err:=_article.Delete()
			//发生错误
			if err!=nil{
				//可能是sql报错
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w,"500服务器内部错误")
			}else {
				if rowsAffected>0{
					indexURL:=route.Name2URL("articles.index")
					http.Redirect(w,r,indexURL,http.StatusFound)
				}else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w,"404文章未找到")
				}
			}
		}
	}
}