package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// Get channel region list
// @router /channel/region [*]
func (this *BaseController) ChannelRegion() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified channel")
		this.ServeJSON()
	}
	num, regions, err := models.GetChannelRegion(channelId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "Not found content")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", regions, num)
	}
	this.ServeJSON()
}

// Get channel type list
// @router /channel/type [*]
func (this *BaseController) ChannelType() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified channel")
		this.ServeJSON()
	}
	num, types, err := models.GetChannelType(channelId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "Not found content")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", types, num)
	}
	this.ServeJSON()
}
