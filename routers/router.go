package routers

import (
	"product/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.ProductController{},"get:ToProduct")
    beego.Router("/login.html", &controllers.LoginController{},"get:ToLogin")
    beego.Router("/login", &controllers.LoginController{},"post:Login")

    beego.Router("/register.html",&controllers.RegisterController{},"get:ToRegister")
    beego.Router("/register",&controllers.RegisterController{},"post:Register")

    beego.Router("/isRegister",&controllers.AJAXController{},"post:IsRegister")

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
