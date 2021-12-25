package entity

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"product/logUtil"
	_ "product/sqlconnect"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 12:15
 **/
type User struct {
	//form 表单    orm 数据库
	Id       int
	Username string `form:"username" orm:"column(username)"`
	Password string `form:"password" orm:"column(password)"`
	Email    string `form:"email" orm:"column(email)"`
}

/*
	定义具体的接口，规定user操作的一系列方法和功能,具体的实现代码在models层
*/
type UserInterface interface {
	Login()
	Register()
	DeleteUser()
	UpdateUser()
	QueryUserByName(name string) (error, User)
	QueryUser(user User) (error, User)
}

//注册模型
/*func init() {
	orm.RegisterModel(new(User))
	orm.RunSyncdb("user_tb", false, true)
}*/
func (u *User) TableName() string {
	return "user_tb"
}

/*
	根据用户名查询，判断用户是否存在，存在返回true，不存在返回false
*/
func (u User) IsRegistered(name string) (err error, bool bool) {
	orm.Debug = true
	ormer := orm.NewOrm()
	ormer.Using("user_tb")
	exist := ormer.QueryTable("user_tb").Filter("username", name).Exist()
	if err != nil {
		logUtil.LogError(err)
		fmt.Println("根据name查询用户失败", err.Error())
		return err, false
	}
	return nil, exist
}

/*
	根据用户信息查询用户，如果查询到了，将user返回，
*/
func (u User) QueryUser(user *User) (u1 User, err error) {
	orm.Debug = true
	ormer := orm.NewOrm()
	ormer.Using("user_tb")

	//方法一：err := ormer.Read(user, "username", "password")
	//注意：ormer.Read(&user,"username"，"password") 查询的时候传参，&用地址传进去的是引用传递。这样会报错，
	//方法二
	err = ormer.QueryTable("user_tb").Filter("username", user.Username).Filter("password", user.Password).One(&u1)
	if err != nil {
		logUtil.LogError(err)
		fmt.Println(err.Error()) //返回no row found
		return u1, err
	}
	return u1, nil
}
