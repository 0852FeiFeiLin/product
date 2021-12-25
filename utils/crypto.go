package utils

import (
	"crypto/md5"
	"fmt"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/18 18:20
 **/
/*
	该函数用于对密码进行加密
*/

func MD5(str string) string {
	md5 := fmt.Sprintf("%X",md5.Sum([]byte(str)))
	fmt.Println(md5)
	return md5
}