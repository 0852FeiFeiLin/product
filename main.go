package main

import (
	"github.com/astaxie/beego/logs"
	_ "product/routers"
	"github.com/astaxie/beego"
)

/*
	设计思路：
		我们具体的模仿是小度商城

	utils层：
		orm操作数据库，增删改查，dbutil层
	entity层：
		写和数据库具体字段对应的实体结构体
	productModel.go层：
		 写和数据库交互的代码，（根据主键id查询商品）
	router层：
		根据具体的请求，分配路由

	productController层：
		写逻辑处理，接收id。调用查询方法，返回详细信息至客户端，

	view层：
		product.html就是商品详情页面，然后如果要购买就提示请先登录
		header.html：页面头部导航栏，固定模块，后面使用哪个template.引入
		然后有一个小度商城的超链接，点击，返回到小度商城，然后可以进行浏览那些一系列的商品，
		我们本地写好几个小队商城里面的具体代表，如果点击查看更多就让他跳转到小度商城里面 https://dumall.baidu.com/


*/
func main() {
	//设置日志 输出格式
	logs.SetLogger("file")
	//设置输出引擎
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/errlogs.log",
	"level":7,
	"maxlines":0,
	"maxsize":0,
	"daily":true,
	"maxdays":10}`)

	beego.Run()
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
}
