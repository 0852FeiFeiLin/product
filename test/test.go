package test

import (
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 11:35
 **/
func main1() {

	//这个是操作数据库了，
	/*	pro := entity.Product{
			Id: 1,
		}
		pro.QueryProduct(pro)
	*/
	//var pduinterface entity.Product
	//QueryProduct(pro)

	//这个只是测试连接数据库，方法里面没有东西
	/*	utils.InsertPro()*/

/*	models.QueryById(1)
*/
/*	user := entity.User{
		Username: "xiaojie",
		Password: utils.MD5("didi"),
		Email: "1820349828@qq.com",
	}
	register := models.Register(&user)
	fmt.Println("受影响的id：",register)
*/

/*	var user entity.User
	err, b := user.IsRegistered("feifei")
	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println("是否存在",b)
	*/
/*
	 user := entity.User{
	 	Username: "linfeifei",
	 	Password: utils.MD5("085200"),
	 }

	queryUser, err := user.QueryUser(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	if queryUser.Id <0 {
		fmt.Println("没有查询到",queryUser.Id)
	}
	fmt.Println(queryUser)*/
	/*
	user := entity.User{
		Username: "linfeifei",
		Password: utils.MD5("085200"),
	}
	deleteUser := models.DeleteUser(&user)
	if deleteUser > 0{
		fmt.Println("删除成功",deleteUser)
	}*/
}
