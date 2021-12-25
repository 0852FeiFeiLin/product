package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"product/entity"
	"product/logUtil"
	"product/models"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/19 16:49
 **/

type RegisterController struct {
	beego.Controller
}

/*
	跳转到注册页面
*/
func (r *RegisterController) ToRegister() {
	r.TplName = "register.html"
}

func (r *RegisterController) Register(){
	var user entity.User
	err := r.ParseForm(&user)
	if err != nil {
		logUtil.LogError(err)
		return
	}
	rePwd := r.GetString("repassword")
	fmt.Println("前端user:",user)
	fmt.Println("前端rePassword:",rePwd)

	//判断两次输入的密码一不一样，不需要，我们在前端阻止了
	rs := models.Register(&user)
	if rs <= 0 {
		fmt.Println("注册失败")
		return
	}
	fmt.Println("注册成功")
	//注册成功就跳转登录
	r.Redirect("/login.html",302)
}