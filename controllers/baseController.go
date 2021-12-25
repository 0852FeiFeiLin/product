package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/28 11:38
 **/
/*
	这是一个基类，相当于类拦截器的那种，
	第一步当然是要在conf里面配置session，
	该类设置一个方法，然后接受者是BaseController，
	然后让别的Controller匿名继承那个BaseController，
	这样就能使用到那个Prepare方法了，然后只需要在登录成功的时候setSession就ok 了
*/

type BaseController struct {
	beego.Controller
	//是否已经登录
	IsLogin bool
	//已经登录的用户信息
	LoginUser interface{}
}
/*
	这是是获取session的方法，如果获取的userName为空，
	那么就代表还没登录，那么就跳转到登录页面
*/
func (c *BaseController)Prepare(){ //这个方法在init之后。发送请求之前
	userName := c.GetSession("userName")
	if userName == nil{
		//若Session中无用户ID则302重定跳转至登陆页面
		//c.Ctx.Redirect(302,"/login.html")
		fmt.Println("请先登录！！！")
		c.IsLogin = false
		c.Data["IsLogin"] = c.IsLogin
		c.Data["userName"] = userName
	}else {//成功登录状态
		c.IsLogin = true
		c.LoginUser = userName
		c.Data["IsLogin"] = c.IsLogin
		c.Data["userName"] = userName
	}
}/*
//session的处理方法2
func (c *BaseController)Prepare()  {
	//获取到已经登录的user
	loginuser := c.GetSession("loginuser")
	if loginuser == nil{
		c.IsLogin = true
		c.Loginuser = loginuser
	}else{
		c.IsLogin = false
	}
	c.Data["IsLogin"] = c.IsLogin
}
*/