package impl

import "product/entity"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/27 23:21
 **/
/*
	定义具体的接口，规定user操作的一系列方法和功能,具体的实现代码在models层
	注意：这一层还只是未完成的，暂时无法解决，但是结构先创在这里了，
*/
/*
	下面的是user结构体特有的方法
*/
type UserInterface interface {
	Login()
	Register()
	DeleteUser()
	UpdateUser()
	QueryUserByName(name string) (error, entity.User)
	QueryUser(user entity.User) (error, entity.User)
}

/*func (u entity.User)Login(){   //会报错
	//我猜测还是不能把他分出来，只能全部写在和结构体同一个的文件中，

}*/