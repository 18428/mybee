package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "mybee/routers"

	_ "github.com/lib/pq"
)

func init() {
	// PostgreSQL 配置
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user=odoo password=odoo dbname=mybee host=127.0.0.1 port=15432 sslmode=disable")

	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	//orm.Debug = true
	//o := orm.NewOrm()
	//o.Using("default")
	//user := new(models.User)
	//user.Name = "tom"
	//user.Age = 25
	//fmt.Println(o.Insert(user))
	beego.Run()
}

