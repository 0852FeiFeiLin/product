package models

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
 * @DateTime: 2021/12/18 10:41
 **/
/*
	orm操作数据库，增删改查
	解耦：userModules ---> dbUtils --->database
	orm已解耦：models ---> orm ---> database  (因为orm就相当于是中间件了)
*/
/*
	查询：传递有数据的结构体过来，然后我们在这里面，根据具体的结构体数据去限定查询条件，然后返回结果集
*/
func QueryById(id int) (entity.Product){
	fmt.Println("根据id进行查询")
	 pro := entity.Product{
	 	Id: id,
	 }
	product := pro.QueryProduct(pro)
	fmt.Println(product)
	return product
}
func QueryResultsById(ids ...int)([]entity.Product,error){
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
	var pros []entity.Product
	//循环参数，然后一次次去请求，然后添加到数组对象里面
	for i:= 0;i<len(ids);i++ {
		 err := ormer.QueryTable("product_tb").Filter("id", ids[i]).One(&product) //存在返回true，不存在返回false
		/*
			用法拓展：
			err := ormer.QueryTable("users").Filter("username", user.Username).Filter("password", user.Password).One(&Users)
			//.one就是返回符合的一条数据，.all就是所有
		*/
		if err != nil {
			//fmt.Println("根据id查询指定产品失败", err.Error())
			//替换成logs日志纪录，写入文件纪录err日志
			logUtil.LogError(err)
			return nil,err
		}
		//添加到数组对象
		pros = append(pros, product)
	}
	fmt.Println(pros)
	return pros,nil
}
/*
	查找精品推荐
*/
func QueryHotPro()( pros []entity.Product,err error){
	/*
		select * from product_tb where id  in(2,5,6);
	*/
	//ids := []int64{2,5,6}
	//返回查询到的产品集
	pros, err = QueryResultsById(2, 5, 6)
	if err != nil {
		logUtil.LogError(err)
		return pros,err
	}
	fmt.Println(pros)
	return pros,nil
}
/*
	产品和数据库交互的代码，查询，
	根据产品名进行查询,并返回查询到的结果
*/
func QueryProByName(name string) (entity.Product){
	fmt.Println("根据name进行查询")
	var pro entity.Product
	product := pro.QueryProductByName(&name)
	fmt.Println(product)
	//返回查询到的产品
	return product
}


/*
	增
*/

/*
	删
*/

/*
	改
*/
