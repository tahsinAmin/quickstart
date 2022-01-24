package routers

import (
	"quickstart/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Show;post:DoShow")
}
