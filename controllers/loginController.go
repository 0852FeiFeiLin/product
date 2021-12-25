package controllers

import (
	"fmt"
	"product/entity"
	"product/logUtil"
	"product/models"
	"product/utils"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/19 16:46
 **/


type LoginController struct {
	BaseController
}
//跳转到登录页面
func (i *LoginController)ToLogin(){
	i.TplName = "login.html"
}
/*
	登录逻辑：解析用户表单，数据传输至数据库进行查询，存在登陆成功，并且设置id为sessionId，.SetSession("userId",Id)
			不存在error,挑战到错误页面
*/
func(i *LoginController) Login(){
	//1、解析前端表单
	var user entity.User
	err := i.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		logUtil.LogError(err)
		return
	}
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	//加密后进行验证
	user.Password = utils.MD5(user.Password)
	fmt.Println(user.Password)
	user1, err := models.Login(&user)
	if err != nil {
		logUtil.LogError(err)
		fmt.Println(err.Error())
		return
	}
	if user1.Id <= 0 {
		i.TplName = "login.html"
		i.Data["errInfo"] = "您输入的用户米和密码错误，请重新输入"
		return
	}
	fmt.Println("登录成功")
	i.SetSession("userName",user1.Username)
	fmt.Println("Session:",i.GetSession("userName"))
	//登录成功跳转product页面
	 i.Data["IsLogin"] = i.IsLogin
	i.Redirect("/product.html",302)
}