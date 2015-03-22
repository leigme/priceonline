package controllers

import (
	"github.com/astaxie/beego"
)

type ManageController struct {
	beego.Controller
}

func (c *ManageController) Get() {
	c.Data["Title"] = "登陆"
	c.TplNames = "admin/signin.html"
}

func (c *ManageController) Admin() {
	c.Data["Title"] = "管理"
	c.News()
}

func (c *ManageController) News() {
	c.Data["NewsActive"] = "active"
	c.Data["Name"] = "资讯管理"
	c.TplNames = "admin/index.html"
}

func (c *ManageController) Addinfo() {
	beego.Info("++++++===++++++")
	texteditor := c.Input().Get("texteditor")
	c.Data["Content"] = texteditor
	c.TplNames = "index.html"
}
