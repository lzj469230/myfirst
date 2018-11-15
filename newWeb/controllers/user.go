package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newWeb/models"
)

type UserController struct{
	beego.Controller
}
//显示注册页面
func (this*UserController)ShowRegister(){
	this.TplName = "register.html"
}
//处理注册业务
func (this*UserController)HandleReg(){
	//接收数据
	userNmae := this.GetString("userName")
	pwd := this.GetString("password")
	//校验数据
	if userNmae == ""||pwd == ""{
		this.Data["errmsg"]="用户名密码不能为空"
		this.TplName = "register.html"
		return
	}
	//插入数据
	o:=orm.NewOrm()
	//插入对象D
	var user models.User
	//给插入对象赋值
	user.UserName = userNmae
	user.Pwd = pwd
	//插入
	_,err := o.Insert(&user)
	if err != nil{
		this.Data["errmsg"] = "注册失败，请重新注册"
		this.TplName = "register.html"
		return
	}
	//返回数据
	this.Redirect("/login",302)
}

//现实登陆页面
func (this*UserController)ShowLogin(){
	this.TplName = "login.html"
}
func(this*UserController)HandleLogin(){
	//接受数据
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	//校验数据
	if userName==""||pwd==""{
		this.Data["errmsg"] = "用户名或在密码不能为空"
		this.TplName = "login.html"
		return
	}

	//处理数据
	//查询业务
	//获取orm对象
	o:=orm.NewOrm()
	//获取查询对象
	var user models.User
	//给查询条件赋值
	user.UserName = userName
	//查询
	err:=o.Read(&user,"UserName")
	if err != nil{
		this.Data["errmsg"] = "用户名不存在"
		this.TplName = "login.html"
		return
	}
	//判断密码是否正确
	if user.Pwd != pwd{
		this.Data["errmsg"] = "密码错误，请重新输入"
		this.TplName = "login.html"
		return
	}
	//返回数据
	this.Redirect("/atticleList",302)
}