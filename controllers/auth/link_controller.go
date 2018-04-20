package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/models"
	"github.com/astaxie/beego"
	"strconv"
	"html/template"
	"github.com/astaxie/beego/validation"
	"fmt"
)

type LinkController struct {
	controllers.BaseController
}

type LinkCreate struct {
	Name 	string 	`form:"name" valid:"Required;MaxSize(23)"`
	Link 	string 	`form:"link" valid:"Required;MaxSize(100)"`
	Sort 	int64   `form:"ordering" valid:"Required"`
}

func (l *LinkController)URLMapping()  {
	l.Mapping("Index",l.Index)
	l.Mapping("Create",l.Create)
	l.Mapping("Store",l.Store)
	l.Mapping("Show",l.Show)
	l.Mapping("Update",l.Update)
	l.Mapping("Destroy",l.Destroy)
}

// @router /console/link [get]
func (l *LinkController) Index() {
	beego.ReadFromRequest(&l.Controller)
	page := l.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	link,_ := models.AllLink()
	totalPage,lastPage,currentPage,nextPage := models.LinkPaginate(page2)
	l.Data["xsrfdata"]= template.HTML(l.XSRFFormHTML())
	l.Data["link"] = link
	l.Data["totalPage"] = totalPage
	l.Data["lastPage"] = lastPage
	l.Data["currentPage"] = currentPage
	l.Data["nextPage"] = nextPage
	l.Layout = "auth/master.tpl"
	l.TplName = "auth/post/index.tpl"

	l.Layout = "auth/master.tpl"
	l.TplName = "auth/link/index.tpl"
}

// @router /console/link/create [get]
func (l *LinkController) Create() {
	beego.ReadFromRequest(&l.Controller)
	l.Data["xsrfdata"]=template.HTML(l.XSRFFormHTML())
	l.Layout = "auth/master.tpl"
	l.TplName = "auth/link/create.tpl"
}

// @router /console/link [post]
func (l *LinkController) Store() {
	u := LinkCreate{}
	valid := validation.Validation{}
	if err := l.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	fmt.Println(u)
	if err != nil {
	}
	if !b {
		_,message :=l.RequestValidate(valid)
		l.MyReminder("error",message)
		l.Redirect("/console/link/create",302)
		return
	}
	link := models.Links{
		Name	:	u.Name,
		Link	:	u.Link,
		Sort	:	u.Sort,
	}
	models.AddLinks(&link)
	l.MyReminder("success","操作成功")
	l.Redirect("/console/link",302)
	return
}

// @router /console/link/:id([0-9]+/edit [get]
func (l *LinkController) Show() {
	beego.ReadFromRequest(&l.Controller)
	id := l.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	link,_ := models.GetLinksById(id64)
	l.Data["link"] = link
	l.Data["xsrfdata"]=template.HTML(l.XSRFFormHTML())
	l.Layout = "auth/master.tpl"
	l.TplName = "auth/link/edit.tpl"
}

// @router /console/link/:id([0-9]+ [put]
func (l *LinkController) Update() {
	u := LinkCreate{}
	valid := validation.Validation{}
	if err := l.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	fmt.Println(u)
	if err != nil {
	}
	if !b {
		_,message :=l.RequestValidate(valid)
		l.MyReminder("error",message)
		l.Redirect("/console/link/create",302)
		return
	}
	id := l.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	link := models.Links{
		Id		:	id64,
		Name	:	u.Name,
		Sort	:	u.Sort,
		Link	:	u.Link,
	}
	models.UpdateLinksById(&link)
	l.MyReminder("success","操作成功")
	l.Redirect("/console/link",302)
	return
}

// @router /console/link/:id([0-9]+ [delete]
func (l *LinkController) Destroy() {
	id := l.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	models.DeleteLinks(id64)
	l.MyReminder("success","操作成功")
	l.Redirect("/console/link",302)
	return
}
