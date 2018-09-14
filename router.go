package routers

import (
	"github.com/switch/firstFrame3/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/video", &controllers.CommonController{}, "post:GetFirstFrame")
}
