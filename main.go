package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "myproject/routers"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:123@/go_test?charset=utf8")
	beego.Run()
}

