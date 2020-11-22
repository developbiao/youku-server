package controllers

import (
	"fyoukuApi/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// User register
// @router /register/save [post]
func (this *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		err      error
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "Mobile can'not is empty")
		this.ServeJSON()
	}
	ok, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !ok {
		this.Data["json"] = ReturnError(4002, "Mobile is not correct")
		this.ServeJSON()
	}
	if password == "" {
		this.Data["json"] = ReturnError(4003, "Password can'not is empty")
		this.ServeJSON()
	}

	// Check mobile has been registered
	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = ReturnError(4005, "This mobile has been registered")
		this.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err != nil {
			this.Data["json"] = ReturnError(5000, err)
		} else {
			this.Data["json"] = ReturnSuccess(0, "Register Success", nil, 0)
		}
		this.ServeJSON()
	}
}

// @router /login/do [*]
func (this *UserController) LoginDo() {
	mobile := this.GetString("mobile")
	password := this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "password can'not is empty")
	}
	ok, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !ok {
		this.Data["json"] = ReturnError(4002, "Mobile is not correct")
		this.ServeJSON()
	}

	if password == "" {
		this.Data["json"] = ReturnError(4003, "Password can'not is empty")
		this.ServeJSON()
	}

	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		this.Data["json"] = ReturnSuccess(0, "Login Success",
			map[string]interface{}{"uid": uid, "username": name}, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "Mobile or password not match")
		this.ServeJSON()
	}
}

// batch send message
// @router /send/message [*]
func (this *UserController) SendMessageDo() {
	uids := this.GetString("uids")
	content := this.GetString("content")

	if uids == "" {
		this.Data["json"] = ReturnError(4001, "请填写接收人~")
		this.ServeJSON()
	}
	if content == "" {
		this.Data["json"] = ReturnError(4002, "请填写发送内容~")
		this.ServeJSON()
	}

	messageId, err := models.SendMessageDo(content)
	if err != nil {
		this.Data["json"] = ReturnError(5000, "发送消息失败，请联系客服~")
		this.ServeJSON()
	}

	uidConfig := strings.Split(uids, ",")
	for _, v := range uidConfig {
		userId, _ := strconv.Atoi(v)
		models.SendMessageUser(int64(userId), messageId)
	}
	this.Data["json"] = ReturnSuccess(0, "发送成功!", "", 1)
	this.ServeJSON()
}
