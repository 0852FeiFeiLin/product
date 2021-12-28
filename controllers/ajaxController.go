package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"product/entity"
	"product/logUtil"
	"product/models"
	"strconv"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/19 19:47
 **/
type AJAXController struct {
	beego.Controller
}

/*
	通过ajax请求，来判断用户名是否已经存在
*/
func (a *AJAXController) IsRegister() {
	ajax := a.IsAjax()
	fmt.Println("是否为ajax请求：", ajax)
	//获取从ajax传过来的数据
	data := a.Ctx.Input.RequestBody
	fmt.Println(data)
	//反序列化
	var user entity.User
	json.Unmarshal(data, &user)
	//把反序列化的用户名传递进行查询
	err, isRegister := user.IsRegistered(user.Username)
	if err != nil {
		fmt.Println(err.Error())
		logUtil.LogError(err)
		return
	}
	//判断是否存在
	if isRegister == false { //不存在
		a.Data["json"] = strconv.FormatBool(isRegister)
		fmt.Println("用户是否已经存在", isRegister)
		a.ServeJSON() //这一步才是发送出去
		return
	}
	a.Data["json"] = strconv.FormatBool(isRegister)
	a.ServeJSON()

	/*
		controller.ServerJSON()自动将模板的数据进行json格式封装，然后显示在前端
		controller.ServerXML() 自动将controller.Data["XML"]中的数据按照XML格式进行组装。
	*/
}

/*
	这个是处理前端用户通过搜索框search发起的请求，
*/
func (a *AJAXController) SearchPro() {
	isajax := a.IsAjax()
	fmt.Println("是否为ajax请求：", isajax)
	//获取从ajax传过来的数据
	data := a.Ctx.Input.RequestBody
	fmt.Println("前端传递到的产品信息：", data)
	//反序列化
	var proName = ""
	json.Unmarshal(data, &proName)
	fmt.Println("前端传递到的产品名称：", proName)
	//调用方法，进行查询，然后f返回指定页面
	pro := models.QueryProByName(proName)
	fmt.Println("根据name查询出来的产品:", pro)
	//我们把产品的id传递给前端，然后前端请求重定向再通过id发起一次请求
	a.Data["json"] = pro.Id
	a.ServeJSON()
}
