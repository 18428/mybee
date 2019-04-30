package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int64
	Name string
	Age  int
	Mark string
}

func init() {
	orm.RegisterModel(new(User))
}