package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

//Create 创建用户，通过User,Id来判断是否成功
func (u *User)Create()(err error)  {
	if err=model.DB.Create(&u).Error;err!=nil {
		logger.LogError(err)
		return err
	}
	return nil
}
//GetByEmail 通过email获取用户
func GetByEmail(email string)(User,error)  {
	var user User
	if err:=model.DB.Where("email=?",email).First(&user).Error;err!=nil{
		return user,err
	}
	return user,nil
}
//Get 通过id获取用户
func Get(idstr string)(User,error)  {
	var user User
	id:=types.StringToInt(idstr)
	if err:=model.DB.First(&user,id).Error;err!=nil{
		return user,err
	}
	return user,nil
}