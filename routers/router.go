package routers

import (
	"quickstart/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Show")
	beego.Router("/ajaxTest", &controllers.MainController{}, "get:AjaxTest")
	beego.Router("/sortData", &controllers.MainController{}, "get:SortData")
	beego.Router("/nextPage", &controllers.MainController{}, "get:NextPage")
}
