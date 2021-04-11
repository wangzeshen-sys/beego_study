package controllers

import (
	"beegodemo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"

	////////////////////////////////////////////////
	// 插入
	// 1 有orm对象
	// o := orm.NewOrm()
	// 2 有一个要插入数据的结构体对象
	// user := models.User{}
	// 3 对结构体对象赋值
	// user.Name = "111"
	// user.Pwd = "222"
	// 4 插入sql
	// _, err := o.Insert(&user)
	// if err != nil {
	// beego.Info("insert failed, err:", err)
	// return
	// }
	c.Data["data"] = "今天中午吃米饭"
	c.TplName = "test.tpl"

	///////////////////////////////////////////////
	// 查询
	// 1 有orm对象
	// o := orm.NewOrm()
	// // 2 查询的对象
	// user := models.User{}
	// // 3 指定查询对象字段值
	// // user.Id = 1
	// user.Name = "111"
	// // 4 查询数据
	// // err := o.Read(&user) // 根据id进行查询
	// err := o.Read(&user, "Name") // 根据name 进行查询
	// if err != nil {
	// 	beego.Info("Read failed, err:", err)
	// 	return
	// }
	// beego.Info("Read success :", user)

	////////////////////////////////////////////////////
	// 更新
	// 1 要有orm对象
	// o := orm.NewOrm()
	// // 2 需要跟新的结构体对象
	// user := models.User{}
	// // 3 查到需要更新的数据
	// user.Id = 1
	// err := o.Read(&user)
	// if err == nil {
	// 	// 4 给数据重新赋值
	// 	user.Name = "222"
	// 	user.Pwd = "333"
	// 	// 5 更新
	// 	_, err := o.Update(&user)
	// 	if err != nil {
	// 		beego.Info("Update failed, err:", err)
	// 		return
	// 	}
	// }


	////////////////////////////////////////
	// 删除
	// 1 有orm 对象
	// o := orm.NewOrm()
	// 2 删除的对象
	// user := models.User{}
	// 3 指定删除那一条对象
	// user.Id = 1
	// 4 删除
	// _, err := o.Delete(&user)
	// if err != nil {
	// 	beego.Info("delete failed, err:", err)
	// 	return
	// }
}

func (c *MainController) Post() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	c.Data["data"] = "今天中午吃面条"
	c.TplName = "test.tpl"
}

func (c *MainController)ShowLogin() {
	c.TplName = "login.tpl"
}

func (c *MainController)HandleLogin(){
	// c.Ctx.WriteString("这是登录的post请求")
	// 1 拿到数据
	userName := c.GetString("username")
	pwd := c.GetString("pwd")
	// 2 判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("数据输入不合法")
		c.TplName = "login.tpl"
		return
	}
	// 3 查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("Read failed, err:", err)
		c.TplName = "login.tpl"
		return 
	}
	if userName != user.Name{
		beego.Info("用户名错误")
		c.TplName = "login.tpl"
		return
	}
	if pwd != user.Pwd {
		beego.Info("密码不正确")
		// 返回返回登录页面的两种方式
		c.TplName = "login.tpl"
		// c.Redirect("/login", 302)
		return
	}
	// 4 跳转
	c.Ctx.WriteString("欢迎您,登录成功")
}














