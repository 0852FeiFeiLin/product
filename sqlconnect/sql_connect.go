package sqlconnect

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"product/logUtil"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 14:49
 **/
func ConnectDB(){//init，在导包的时候就自动连接上了数据库
	//读取配置文件
	config := beego.AppConfig
	//key -->value 根据key找到匹配的值
	//数据库驱动
	driverName := config.String("db_driverName")
	//用户名
	user := config.String("db_user")
	//密码
	pwd := config.String("db_pwd")
	//连接ip
	ip := config.String("db_ip")
	//端口 3306端口
	port := config.String("db_port")
	//数据库名
	db_name := config.String("db_name")

	//注册驱动
	err := orm.RegisterDriver(driverName,  orm.DRMySQL)
	if err != nil {
		//错误写入日志中
		logUtil.LogError(err)
		fmt.Println("注册驱动失败。。。。",err.Error())
		return
	}
	//连接数据库
	/*
		参数1：数据库别名
		参数2：驱动名称
		参数3：数据库信息
	*/
	err = orm.RegisterDataBase("default", driverName, user+":"+pwd+"@tcp("+ip+":"+port+")/"+db_name+"?charset=utf8")
	if err != nil {
		logUtil.LogError(err)
		fmt.Println("连接数据库失败。。。。",err.Error())
		return
	}
	fmt.Println("product_db连接成功！")
}