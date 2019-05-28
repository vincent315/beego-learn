package routers

import (
	"github.com/astaxie/beego"
	"myproject/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
    //beego.Router("/test", &controllers.TestController{})
    beego.Include(&controllers.TestController{})
}
