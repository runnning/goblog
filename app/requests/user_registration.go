package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/user"
)

func ValidateRegistrationForm(data user.User)map[string][]string  {
	//定制认证规则
	rules:=govalidator.MapData{
		"name":[]string{"required","alpha_num","between:3,20","not_exists:users,name"},
		"email":[]string{"required","min:4","max:30","email","not_exists:users,email"},
		"password":[]string{"required","min:6"},
		"password_confirm":[]string{"required"},
	}
	//定制错误信息
	messages:=govalidator.MapData{
		"name":[]string{
			"required:用户名为必选项",
			"alpha_num:格式错误,只允许数字和英文",
			"between:用户名长度需要在3~20之间",
		},
		"email":[]string{
			"required:email为必选项",
			"min:email长度大于4",
			"max:email长度小于30",
			"email:email格式不正确,请提供有效的邮箱地址",
		},
		"password":[]string{
			"required:密码为必填项",
			"min:长度需大于6",
		},
		"password_confirm":[]string{
			"required:确认密码框为必填项",
		},
	}
	//配置初始化
	opts:=govalidator.Options{
		Data: &data,
		Rules: rules,
		TagIdentifier: "valid",//模型中的struct 标签标识符
		Messages: messages,
	}
	//开始验证
	errs:=govalidator.New(opts).ValidateStruct()
	//自定义password_confirm验证
	if data.Password!=data.PasswordConfirm {
		errs["password_confirm"]=append(errs["password_confirm"],"两次输入密码不匹配!")
	}
	return errs
}
