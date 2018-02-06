package template

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	"github.com/wang22/zing/src/controller"
	"github.com/wang22/zing/src/util"
)

// macro 匹配define
var macroRegexp = regexp.MustCompile(`\{\{template \"(.+?)\"[\s\\.]*?\}\}`)

// Renderer 模板渲染
type Renderer struct {
	CacheMap map[string]*template.Template
}

// Render 渲染方法
func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	ctx := data.(*controller.HTTPContext)
	if cache, exists := r.checkCache(name); exists {
		return r.execute(cache, ctx)
	}
	var tmpl *template.Template
	layoutInclude := util.NewSet()
	if ctx.Router.Layout != "" { // 如果有布局文件，则创建布局模板
		tpl := template.New(name)
		content := r.readTemplate(ctx.Router.Layout)
		tmpl, _ = tpl.Parse(content)
		r.findInclude(layoutInclude, content)
	} else {
		tpl := template.New(name)
		tmpl, _ = tpl.Parse(`{{template "content" .}}`)
	}

	// 获取macro文件
	files := util.FileUtil().FileList(util.FileUtil().AbsPath(ctx.Router.TemplateDir), true)
	var macro *template.Template
	if len(files) > 0 {
		var err error
		macro, err = template.ParseFiles(files...)
		if err != nil {
			return err
		}
	}

	tpl := template.New(name + "_view")
	view, _ := tpl.Parse(r.readTemplate(name))
	layoutInclude.Each(func(v interface{}) {
		name := v.(string)
		if t := view.Lookup(name); t != nil {
			tmpl.AddParseTree(name, view.Lookup(name).Tree)
		} else if macro != nil {
			if t := macro.Lookup(name); t != nil {
				tmpl.AddParseTree(name, macro.Lookup(name).Tree)
			}
		} else {
			tmpl.Parse(`{{define "` + name + `"}}{{end}}`)
		}
	})

	// 设置缓存
	if util.Config().TemplateCache {
		r.CacheMap[name] = tmpl
	}

	return r.execute(tmpl, ctx)
}

func (r *Renderer) readTemplate(file string) string {
	absPath := util.FileUtil().AbsPath(file)
	return util.FileUtil().ReadFile(absPath)
}

func (r *Renderer) findInclude(include *util.Set, content string) {
	matchs := macroRegexp.FindAllStringSubmatch(content, -1)
	if len(matchs) > 0 {
		for _, val := range matchs {
			include.Add(val[1])
		}
	}
}

func (r *Renderer) execute(tmpl *template.Template, ctx *controller.HTTPContext) error {
	// 开始渲染
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, ctx.Data)
	if err != nil {
		util.Log().Error("执行页面渲染失败")
		util.Log().Error(err)
		return err
	}
	return ctx.Context.HTML(http.StatusOK, buf.String())
}

func (r *Renderer) checkCache(name string) (*template.Template, bool) {
	if !util.Config().TemplateCache {
		return nil, false
	}
	tmpl := r.CacheMap[name]
	if tmpl == nil {
		return nil, false
	}
	return tmpl, true
}
