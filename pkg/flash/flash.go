package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

//Flashes Flash 消息数组类型，用以在会话中存储map
type Flashes map[string]interface{}

//存入会话数据里的key
var flashkey="_flashes"

func init()  {
	//在 gorilla/sessions 中存储 map 和 struct 数据
	//提前注册gob,方便后续gob序列化编码、解码
	gob.Register(Flashes{})
}

//Info 添加Info类型的消息提示
func Info(message string)  {
	addFlash("info",message)
}

//Warning 添加Warning类型的消息提示
func Warning(message string)  {
	addFlash("warning",message)
}

//Success 添加Success类型的消息提示
func Success(message string)  {
	addFlash("success",message)
}

//Danger 添加Danger类型的消息提示
func Danger(message string)  {
	addFlash("danger",message)
}

func All()Flashes  {
	val:=session.Get(flashkey)
	//类型检测
	flashMessages,ok:=val.(Flashes)
	if!ok{
		return nil
	}
	//读取即销毁,直接删除
	session.Forget(flashkey)
	return flashMessages
}

//私有方法,新增一条提示
func addFlash(key string,message string)  {
	flashes:=Flashes{}
	flashes[key]=message
	session.Put(flashkey,flashes)
	session.Save()
}