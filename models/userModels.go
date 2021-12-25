package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"product/entity"
	"product/logUtil"
	"product/utils"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 10:40
 **/

/*
	用户和数据库交互的代码：注册，登录，更新，注销等等
*/
/*
	登录：根据用户时输入的用户名和密码进行查询，密码在传递前进行md4加密了
		登录成功并返回查询到的用户
*/
func Login(user *entity.User)(entity.User,error){
	//开启orm调试
	//根据值进行查询
	queryUser, err := user.QueryUser(user)
	if err != nil {
		logUtil.LogError(err)
		fmt.Println(err.Error())
	}
	if queryUser.Id <0 {
		fmt.Println("没有查询到",queryUser.Id)
	}
	fmt.Println(queryUser)
	//返回查询到的user
	return queryUser,nil
}
/*
	插入：根据传递过来的结构体插入用户，并返回受影响的行数
*/
func Register(user *entity.User) int64 {
	orm.Debug = true
	ormer := orm.NewOrm()
	ormer.Using("user_tb")
	//先判断用户名是否存在，存在不能插入
	var user1 entity.User
	err, exist := user1.IsRegistered(user.Username)
	fmt.Println("用户是否已经存在：",exist)
	if err != nil {
		logUtil.LogError(err)
		fmt.Println(err.Error(),"注册失败")
		return -1
	}
	if exist ==true{
		fmt.Println("用户名已经存在")
		return -1
	}
	user.Password = utils.MD5(user.Password)
	//Insert 方法的参数必须是指针，返回值是插入数据的主键
	insert, err := ormer.Insert(user)
	if err != nil {
		logUtil.LogError(err)
		fmt.Println(err.Error())
		return -1
	}
	return insert
}


/*
	注销：根据用户输入的用户名和密码，进行delete操作
*/

func DeleteUser(user *entity.User)(int64){
	//打开orm调试
	orm.Debug = true
	//创建orm对象
	ormer := orm.NewOrm()
	//切换到另一个数据库驱动程序
	ormer.Using("user_tb")
	delete, err := ormer.Delete(user, "username", "password") //字段作为查询条件，查找纪录，然后删除
	if err != nil {
		logUtil.LogError(err)
		fmt.Println(err.Error())
		return -1
	}
	return delete
}