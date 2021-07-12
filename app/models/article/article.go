package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

//Article 文章模型
type Article struct {
	models.BaseModel
	Title string `valid:"title"`
	Body string`valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	CategoryID uint64 `gorm:"not null;default:3;index"`
	User user.User
}

//Link 方法用来生产文章链接
func (a Article)Link() string  {
	return route.Name2URL("articles.show","id",strconv.FormatInt(int64(a.ID),10))
}
func (a Article)CreatedAtDate()string  {
	return a.CreatedAt.Format("2006-01-02")
}
