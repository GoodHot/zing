package controller

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//////////////////////////////////////////////////////////////////
////////////////////////// Http Context //////////////////////////
//////////////////////////////////////////////////////////////////

// NewHTTPContext 创建新的上下文
func NewHTTPContext(c echo.Context) *HTTPContext {
	return &HTTPContext{
		Context: c,
		Data:    make(map[string]interface{}),
	}
}

// HTTPContext Http上下文
type HTTPContext struct {
	Context echo.Context           // echo 上下文
	Data    map[string]interface{} // 数据载体
	Router  *HTTPRouter            // 路由器
}

// Put 设置数据
func (ctx *HTTPContext) Put(key string, value interface{}) {
	ctx.Data[key] = value
}

// Param 获取参数
func (ctx *HTTPContext) Param(key string) string {
	return ctx.Context.Param(key)
}

// ParamInt 获取Int参数
func (ctx *HTTPContext) ParamInt(key string) int {
	value, err := strconv.Atoi(ctx.Param(key))
	if err != nil {
		return -1
	}
	return value
}

// QueryParam 获取参数
func (ctx *HTTPContext) QueryParam(key string) string {
	return ctx.Context.QueryParam(key)
}

// QueryParamInt 获取Int参数
func (ctx *HTTPContext) QueryParamInt(key string) int {
	value, err := strconv.Atoi(ctx.QueryParam(key))
	if err != nil {
		return -1
	}
	return value
}

// Bind 绑定参数
func (ctx *HTTPContext) Bind(i interface{}) error {
	return ctx.Context.Bind(i)
}

// FormFile 表单文件
func (ctx *HTTPContext) FormFile(name string) (*multipart.FileHeader, error) {
	return ctx.Context.FormFile(name)
}

// JSON JSON渲染
func (ctx *HTTPContext) JSON() error {
	return ctx.Context.JSON(http.StatusOK, ctx.Data)
}

// Render 模板渲染
func (ctx *HTTPContext) Render(view string) error {
	return ctx.Context.Render(http.StatusOK, ctx.Router.TemplateDir+"/"+view+".html", ctx)
}

// Response echo Response
func (ctx *HTTPContext) Response() *echo.Response {
	return ctx.Context.Response()
}

// Request echo Response
func (ctx *HTTPContext) Request() *http.Request {
	return ctx.Context.Request()
}

// IP 获取客户端IP
func (ctx *HTTPContext) IP() string {
	return ctx.Context.RealIP()
}

// SetCookie 设置Cookie
func (ctx *HTTPContext) SetCookie(cookie *http.Cookie) {
	ctx.Context.SetCookie(cookie)
}

// Cookie 获取Cookie
func (ctx *HTTPContext) Cookie(name string) (*http.Cookie, error) {
	return ctx.Context.Cookie(name)
}

// Redirect 重定向
func (ctx *HTTPContext) Redirect(uri string) error {
	return ctx.Context.Redirect(http.StatusFound, uri)
}

//////////////////////////////////////////////////////////////////
////////////////////////// Http router ///////////////////////////
//////////////////////////////////////////////////////////////////

// HTTPRouter Http路由器
type HTTPRouter struct {
	Group       *echo.Group
	Layout      string // 布局模板
	TemplateDir string // 模板文件夹
}

// Get get请求
func (router *HTTPRouter) Get(uri string, f func(ctx *HTTPContext) error) {
	router.Layout = router.TemplateDir + router.Layout
	router.Group.GET(uri, func(c echo.Context) error {
		context := NewHTTPContext(c)
		context.Router = router
		return f(context)
	})
}

// Post get请求
func (router *HTTPRouter) Post(uri string, f func(ctx *HTTPContext) error) {
	router.Layout = router.TemplateDir + router.Layout
	router.Group.POST(uri, func(c echo.Context) error {
		context := NewHTTPContext(c)
		context.Router = router
		return f(context)
	})
}
