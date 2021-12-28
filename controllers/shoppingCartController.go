package controllers

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/24 18:04
 **/
type ShoppingCartController struct {
	BaseController
}
func (s *ShoppingCartController)ToCart(){
	s.TplName = "shoppingCart.html"
}