package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})

	// 注意：当实现了自定义的get请求方法，请求将不会访问默认方法
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin") // 指定 get方法去实现ShowLogin的方法
}
