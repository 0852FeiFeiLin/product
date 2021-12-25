package entity

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"product/logUtil"
	"product/sqlconnect"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 10:27
 **/
/*
	本类对应的是数据库产品的实体类，实现具体字段
*/
type Product struct {
	Id          int    `form:"id" orm:"column(id)"`
	ProductName string `form:"productName" orm:"column(product_name)"`
	Producer    string `producer:"producer" orm:"column(producer)"`
	CreateTime  string `form:"createTime" orm:"column(createTime)"`
	Content string `form:"content" orm:"column(content)"`
	Price float64 `form:"price" orm:"column(price)"`
}

type ProductInterface interface {
	QueryProduct(product Product) Product
	QueryProductByName(product Product) Product
}

func init() {
	//连接数据库
	sqlconnect.ConnectDB()
	orm.RegisterModel(new(Product))//模型1
	orm.RegisterModel(new(User)) //模型2
	orm.RunSyncdb("default", false, true)
}

//自定义表名，(系统会自动调用)
func (product *Product) TableName() string {
	return "product_tb"
}
/*
	查询产品的方法实现，具体定义在Product接口里面
*/
func (p Product) QueryProduct(pro Product) Product {
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
	fmt.Println("根据id查询指定产品")
	var product Product
	err := ormer.QueryTable("product_tb").Filter("id", pro.Id).One(&product) //存在返回true，不存在返回false
	/*
		用法拓展：
		err := ormer.QueryTable("users").Filter("username", user.Username).Filter("password", user.Password).One(&Users)
		//.one就是返回符合的一条数据，.all就是所有
	*/
	if err != nil {
		fmt.Println("根据id查询指定产品失败", err.Error()) //因为我们还在测试，所以，先不删，后期项目写完了，就删掉这一行。
		logUtil.LogError(err)
		return Product{}
	}
	return product
}
func (p Product) QueryProductByName(proName *string) (pro Product){
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
	fmt.Println("根据name查询指定产品:",proName)
	var product1 Product
	//product_name__contains 是否包含
	err := ormer.QueryTable("product_tb").Filter("product_name", proName).One(&product1) //存在返回true，不存在返回false
	if err != nil {
		fmt.Println("根据id查询指定产品失败", err.Error()) //因为我们还在测试，所以，先不删，后期项目写完了，就删掉这一行。
		logUtil.LogError(err)
		return Product{}
	}
	return product1
}