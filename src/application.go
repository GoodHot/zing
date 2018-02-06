package src

import (
	"fmt"
	"html/template"

	"github.com/labstack/echo"
	"github.com/wang22/zing/src/controller/admin"
	tmpl "github.com/wang22/zing/src/template"
	"github.com/wang22/zing/src/util"
)

// ServerStart 启动服务器
func ServerStart() {
	util.InitialLog("config/seelog.xml")
	util.InitialConfig("config/config.json")

	e := echo.New()
	e.Static("/static", "static")
	e.Renderer = &tmpl.Renderer{
		CacheMap: make(map[string]*template.Template),
	}

	admin.AdminRouter(e.Group(util.Config().AdminURI))

	err := e.Start(util.Config().Port)

	if err != nil {
		fmt.Println("Start error...")
		fmt.Println(err.Error())
	}
}
