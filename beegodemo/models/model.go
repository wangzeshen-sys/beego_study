package models

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

// model 中放置 表的设计
type User struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	// 设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:asdasd123@tcp(127.0.0.1:3306)/db3?charset=utf8")  // 第一个参数  别名，不常用,一般使用default
	// 映射model 数据 对应结构体的名称
	orm.RegisterModel(new(User))
	// 生成表
	orm.RunSyncdb("default", false, true)
	// 第一个参数：数据库别名
	// 第二个参数：是否强制更新，默认不更新
	// 第三个参数：创建的表的sql语句是否显示在终端
}

