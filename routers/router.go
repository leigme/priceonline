package routers

import (
	"github.com/astaxie/beego"
	"priceonline/controllers"
)

func init() {
	beego.Router("/test", &controllers.MainController{}, "get:Test")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/getajaxarticles", &controllers.MainController{}, "get:GetAjaxArticles")
	beego.Router("/topic", &controllers.MainController{}, "get:Topic")
	beego.Router("/policies", &controllers.MainController{}, "get:GetPolicies")

	beego.Router("/admin/signin", &controllers.ManageController{})
	beego.Router("/admin", &controllers.ManageController{}, "post:Admin")
	beego.Router("/admin/addinfo", &controllers.ManageController{}, "post:Addinfo")
}
