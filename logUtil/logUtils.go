package logUtil

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/20 16:39
 **/
/*
	这个util就是具体的日志处理方法
*/
func LogInfo(v interface{}){ //v 就是我们要传递的日志
	conf := beego.AppConfig
	//拿到运行模式
	runMode := conf.String("runmode")
	if runMode == "dev"{
		logs.Info(v)  //写入日志
	}
}
func LogError(v interface{}){
	conf := beego.AppConfig
	//拿到运行模式
	runMode := conf.String("runmode")
	if runMode == "dev"{
		logs.Error(v)  //写入日志
	}
}
func LogWarning(v interface{}){
	conf := beego.AppConfig
	//拿到运行模式
	runMode := conf.String("runmode")
	if runMode == "dev"{
		logs.Warning(v)  //写入日志
	}
}
func LogDebug(v interface{}){
	conf := beego.AppConfig
	//拿到运行模式
	runMode := conf.String("runmode")
	if runMode == "dev"{
		logs.Debug(v)  //写入日志
	}
}

func LogTrace (v interface{}){
	conf := beego.AppConfig
	//拿到运行模式
	runMode := conf.String("runmode")
	if runMode == "dev"{
		logs.Trace(v)  //写入日志
	}
}

/*
	然后在实际的日志处理中，根据具体的类型，使用这些方法代替fmt.Println()
	例如正常：LogInfo， err错误：LogError
	如果要写入文件，先创建一个.log文件，然后配置相应的参数在main里面。
*/

