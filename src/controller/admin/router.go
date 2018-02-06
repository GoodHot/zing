package admin

import (
	"github.com/labstack/echo"
	"github.com/wang22/zing/src/controller"
	"github.com/wang22/zing/src/util"
)

// AdminRouter admin路由
func AdminRouter(group *echo.Group) {
	router := new(controller.HTTPRouter)
	router.Group = group
	router.TemplateDir = util.Config().TemplateDir + "/admin/"
	router.Layout = "layout.html"

	index := new(IndexController)
	router.Get("", index.Index)
}
