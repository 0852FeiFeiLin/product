package controllers

import (
	"fmt"
	"product/logUtil"
	"product/models"
	"product/utils"
	"strconv"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/19 23:29
 **/

type ProductInfoController struct {
	BaseController
}

//直接点击一个产品跳转到相应的内容页面，不需要刻意的跳转产品详情页面，刻意跳转，也不知道是哪个产品的信息啊
func (p * ProductInfoController)ToProductInfo()  {
	/*p.TplName = "productInfo.html"
	p.Data["productInfo"] = "124345"
	p.Redirect("/productInfo",302)*/
}
func (p *ProductInfoController) ProductInfo() {
	//获取到id
	valueID := p.Input().Get("id")
	fmt.Println(valueID)
	producer := p.GetString("producer")
	price := p.GetString("price")
	fmt.Println("价格：",price)
	fmt.Println("producer值:",producer)

	if producer == "京东"{
		fmt.Println("我是京东商品")
		p.TplName = "productInfo2.html" //京东商城
	}else{
		p.TplName = "productInfo.html" //小度商城
	}
	//id转换为int传递
	id,err := strconv.Atoi(valueID)
	fmt.Println(id)
	if err != nil {
		logUtil.LogError(err)
		return
	}
	//根据id查询出来商品
	productInfo := models.QueryById(id)
	fmt.Println(productInfo)
	//将时间转换成标准时间格式输出
	createTime, err := strconv.ParseInt(productInfo.CreateTime, 10, 64)
	if err != nil {
		logUtil.LogError(err)
		return
	}
	productInfo.CreateTime = utils.SwitchTimeStampToData(createTime)
	//加上价格 (不需要了，直接在数据库里面加了字段，之前没在数据库里面加字段，就是通过url请求后面接的price=9.9)
	/*productInfo.Price = price*/
	//转换为map发送给前端的模板
	pros := map[string]interface{}{
		"ProductInfo" : productInfo,
	}
	p.Data["Price"] = productInfo.Price //这个是那个大价格
	p.Data["ProductInfo"] = pros //前端{{range ProductInfo}}

}
