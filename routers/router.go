package routers

import (
	"product/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.ProductController{},"get:ToProduct")//产品页

    beego.Router("/login.html", &controllers.LoginController{},"get:ToLogin")
    beego.Router("/login", &controllers.LoginController{},"post:Login")

    beego.Router("/register.html",&controllers.RegisterController{},"get:ToRegister")
    beego.Router("/register",&controllers.RegisterController{},"post:Register")

    //利用ajax判断是否注册过了
    beego.Router("/isRegister",&controllers.AJAXController{},"post:IsRegister")

    //显示和产品相关
    beego.Router("/product.html",&controllers.ProductController{},"get:ToProduct")

    //产品详情页面
	beego.Router("/productInfo",&controllers.ProductInfoController{},"get:ProductInfo")

    //用户根据搜索框搜索
	beego.Router("/searchPro",&controllers.AJAXController{},"post:SearchPro")

    //退出
    beego.Router("/exit",&controllers.ExitController{},"get:Exit")

    //购物车
    beego.Router("/shoppingCart.html",&controllers.ShoppingCartController{},"get:ToCart")


}
