package auth

import (
	"bee-go-myBlog/controllers"
	"bee-go-myBlog/services"
	"strconv"
	"bee-go-myBlog/models"
	"github.com/astaxie/beego"
	"html/template"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/orm"
)

type TagController struct {
	controllers.BaseController
}

type TagRequest struct {
	Name 	string `form:"name" valid:"Required;MaxSize(30)"`
}

func (t *TagController)URLMapping()  {
	t.Mapping("Index",t.Index)
	t.Mapping("Create",t.Create)
	t.Mapping("Store",t.Store)
	t.Mapping("Show",t.Show)
	t.Mapping("Update",t.Update)
	t.Mapping("Destroy",t.Destroy)
	t.Mapping("GetTagByLike",t.GetTagByLike)
}


// @router /console/tag [get]
func (t *TagController) Index() {
	beego.ReadFromRequest(&t.Controller)
	page := t.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	tag,_ := services.GetAllMyTag(page2)

	totalPage,lastPage,currentPage,nextPage := models.TagPaginate(page2)
	t.Data["xsrfdata"]=template.HTML(t.XSRFFormHTML())
	t.Data["totalPage"] = totalPage
	t.Data["lastPage"] = lastPage
	t.Data["currentPage"] = currentPage
	t.Data["nextPage"] = nextPage
	t.Data["tag"] = tag
	t.Layout = "auth/master.tpl"
	t.TplName = "auth/tag/index.tpl"
}

// @router /console/tag/create [get]
func (t *TagController) Create() {
	beego.ReadFromRequest(&t.Controller)
	t.Data["xsrfdata"]=template.HTML(t.XSRFFormHTML())
	t.Layout = "auth/master.tpl"
	t.TplName = "auth/tag/create.tpl"
}

// @router /console/tag [post]
func (t *TagController) Store() {
	u := TagRequest{}
	valid := validation.Validation{}
	if err := t.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	if err != nil {
	}
	if !b {
		_,message := t.RequestValidate(valid)
		t.MyReminder("error",message)
		t.Redirect("/console/tag/create",302)
		return
	}
	var tagCreate = models.Tags{
		Name	:	u.Name,
	}
	_,err = models.AddTags(&tagCreate)
	if err != nil {
		t.MyReminder("error","系统内部错误")
	}
	t.Redirect("/console/tag",302)
}

// @router /console/tag/:id([0-9]+/edit [get]
func (t *TagController) Show() {
	beego.ReadFromRequest(&t.Controller)
	id := t.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	tag,_ := models.GetTagsById(id64)
	t.Data["xsrfdata"]=template.HTML(t.XSRFFormHTML())
	t.Data["tag"] = tag
	t.Layout = "auth/master.tpl"
	t.TplName = "auth/tag/edit.tpl"
}

// @router /console/tag/:id([0-9]+ [put]
func (t *TagController) Update() {
	u := TagRequest{}
	valid := validation.Validation{}
	if err := t.ParseForm(&u); err != nil {
		fmt.Println(err)
	}
	b, err := valid.Valid(&u)
	if err != nil {
	}
	if !b {
		_,message := t.RequestValidate(valid)
		t.MyReminder("error",message)
		t.Redirect("/console/tag/create",302)
		return
	}
	id := t.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	o := orm.NewOrm()
	tag := models.Tags{Id:id64}
	err = o.Read(&tag)
	if err != nil {
		t.MyReminder("error","数据不存在,请检查后再试")
		t.Redirect("/console/tag",302)
		return
	}
	if u.Name == tag.Name {
		t.MyReminder("error","未做任何修改")
		t.Redirect("/console/tag/"+id+"/edit",302)
		return
	}

	searchTag := models.Tags{}
	num,_ := o.QueryTable(new(models.Tags)).Filter("Name",u.Name).Filter("Id__ne",id64).All(&searchTag)
	if num > 0 {
		t.MyReminder("error","标签名已存在,请检查后再试")
		t.Redirect("/console/tag/"+id+"/edit",302)
		return
	}
	var tagUpdate models.Tags
	tagUpdate.Id = id64
	tagUpdate.Name = u.Name
	o.Update(&tagUpdate)
	t.MyReminder("success","操作成功")
	t.Redirect("/console/tag",302)
	return
}

// @router /console/tag/:id([0-9]+ [delete]
func (t *TagController) Destroy() {
	id := t.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	res := services.DeleteTag(id64)
	if res {
		t.MyReminder("success","操作成功")
		t.Redirect("/console/tag",302)
	} else {
		t.MyReminder("error","操作失败")
		t.Redirect("/console/tag",302)
	}
}

// @router /console/tag/auto [get]
func (t *TagController) GetTagByLike() {
	param := t.GetString("term")
	res := services.GetTagByLike(param)
	t.Data["json"] = res
	t.ServeJSON()
}