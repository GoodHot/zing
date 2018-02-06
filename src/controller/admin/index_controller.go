package admin

import (
	"github.com/wang22/zing/src/controller"
)

// IndexController 首页
type IndexController struct {
}

// Index 首页
func (IndexController) Index(ctx *controller.HTTPContext) error {
	ctx.Put("hello", "world")
	return ctx.Render("index")
}
