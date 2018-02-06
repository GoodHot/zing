package util

import (
	"encoding/json"
	"net/http"
	"time"
)

// Config 获取到当前环境的配置
func Config() *Configure {
	return cfg
}

var cfg = &Configure{}

// InitialConfig 配置初始化
func InitialConfig(cfgPath string) {
	path := FileUtil().AbsPath(cfgPath)
	cfgFile := FileUtil().ReadFile(path)
	json.Unmarshal([]byte(cfgFile), cfg)
}

// Configure 环境配置
type Configure struct {
	Port          string `json:"port"`          // 端口
	AdminURI      string `json:"adminURI"`      // Admin访问URI
	TemplateDir   string `json:"templateDir"`   // 模板文件目录
	TemplateCache bool   `json:"TemplateCache"` // 模板是否缓存
	Cookie        struct {
		Domain   string `json:"domain"`   // cookie 域名
		Path     string `json:"path"`     // cookie 路径
		HTTPOnly bool   `json:"httpOnly"` // cookie HttpOnly
		Secure   bool   `json:"secure"`   // cookie Secure
	} `json:"cookie"` // cookie配置
}

// NewCookie 创建新Cookie
func (cfg Configure) NewCookie(name string, val string, expires time.Time) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Domain = cfg.Cookie.Domain
	cookie.Path = cfg.Cookie.Path
	cookie.HttpOnly = cfg.Cookie.HTTPOnly
	cookie.Secure = cfg.Cookie.Secure
	cookie.Name = name
	cookie.Value = val
	cookie.Expires = expires

	return cookie
}
