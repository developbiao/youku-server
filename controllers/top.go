package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

// Get ranking by channel
// @router /channel/top [*]
func (this *TopController) ChannelTop() {
	// Get channel id
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelTop(channelId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	}
	this.ServeJSON()
}

// Get ranking by type
// @router /type/top [*]
func (this *TopController) TypeTop() {
	typeId, _ := this.GetInt("typeId")
	if typeId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定类型")
		this.ServeJSON()
	}
	num, videos, err := models.GetTypeTop(typeId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有找到相关内容")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	}
	this.ServeJSON()
}
