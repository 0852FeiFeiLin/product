package controllers

import "fmt"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/19 22:44
 **/
type ExitController struct{
	BaseController
}
func (this ExitController)Exit(){
	//这里退出应该是返回到产品详情页面

	//删除session
	this.DelSession("userName")
	fmt.Println("Session:",this.GetSession("userName"))//检查

	//请求重定向
	this.Redirect("/",302)
}