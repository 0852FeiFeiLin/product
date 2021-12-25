package utils

import "time"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/1 20:02
 **/
/*
	将传入的时间戳转换为时间
*/
func SwitchTimeStampToData(timeStamp int64)string{
	//时间戳转换为时间格式
	t := time.Unix(timeStamp,0)
	//格式化为标准时间格式
	return t.Format("2006-01-02 15:04:05")
}

