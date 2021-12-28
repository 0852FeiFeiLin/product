package controllers

import (
	"fmt"
	"product/models"
	"strconv"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 10:35
 **/
/*
	这个是product产品的控制层，根据具体的请求，来实现具体的操作
*/
type ProductController struct {
	BaseController
}
/*
	跳转到product.html
*/
func (p *ProductController) ToProduct() {
	p.TplName = "product.html"
	p.Data["IsLogin"] = p.IsLogin

	/*去数据库查找，然后显示在模板*/
	/*	p.Redirect("/product",302)
	 */
	//精品推荐那里显示的产品
	pros, err := models.QueryHotPro()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//利用模板语法然后把产品集里面的product_name 渲染到前端，还有id隐藏渲染
	//创建map
	var products = make(map[string]interface{})
	for i, pro := range pros {
		//全部添加到键值对的map集合里
		products["Product"+strconv.Itoa(i)] = pro
		/*p.Data["Product"+strconv.Itoa(i)] = pro.ProductName*/
	}
	fmt.Println(products)
	//遍历map，然后将map的值全部模板语法发送到前端
	for k, v := range products {
	/*	fmt.Println(k )
		fmt.Println(v)*/
		p.Data[k] = v   //p.Data["Product1"] = pro.ProductName
	}
	/*p.Data["Product0"] = products["Product0"]
	p.Data["Product1"] = products["Product1"]*/
}

