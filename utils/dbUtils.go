package utils

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"product/entity"
	"product/logUtil"
)
/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 10:26
 **/
/*
	orm操作数据库，增删改查
	解耦：userModules ---> dbUtils --->database
	models ---> orm ---> database
*/

/*
	连接 ，好像这一系列操作都能直接写在models层
*/

/*
	查询：传递有数据的结构体过来，然后我们在这里面，根据具体的结构体数据去限定查询条件，然后返回结果集
*/
func QueryProduct(pro entity.Product)(p entity.Product){
	//开启可调式
	orm.Debug = true
	//新建orm对象
	ormer := orm.NewOrm()
	//切换到default程序
	ormer.Using("default")
	//调用查询方法
	/*
		QueryTable 查询到表所有数据
		Filter 根据条件过滤
		Exist 是否存在，返回bool
	*/
	var product entity.Product
	err := ormer.QueryTable("product_tb").Filter("id", pro.Id).One(&product)  //存在返回true，不存在返回false
/*
	用法拓展：
	err := ormer.QueryTable("users").Filter("username", user.Username).Filter("password", user.Password).One(&Users)
	//.one就是返回符合的一条数据，.all就是所有
	*/
	if err != nil {
		logUtil.LogError(err)
		fmt.Println("根据id查询指定产品失败",err.Error())
		return
	}
	return product
}
/*
	增
*/
func InsertPro(){
	fmt.Println(11111111111111)

}
/*
	删
*/

/*
	改
*/
