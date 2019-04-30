package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mybee/models"
)

type EditorController struct {
	beego.Controller
}

func (this *EditorController) Get()  {
	this.Data["name"] = "yanhuaqiang"
	this.Data["age"] = 29
	this.Data["sex"] = "男"

	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "Paris"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	this.Data["country"] = countryCapitalMap



	o := orm.NewOrm()
	user := models.User{Id: 1}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name, user.Mark)
	}


	this.TplName = "editor.html"
}