package controllers

import (
	"fyoukuApi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// User register
// @router /register/save [get]
func (this *UserController) SaveRegister() {
	var(
		mobile 		string
		password 	string
		err 		error
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.Data["json"] = 0
	}
}



