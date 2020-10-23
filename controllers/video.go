package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// Channel page - get channel advert
// @router /channel/advert [*]
func (this *VideoController) ChannelAdvert() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified channel")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelAdvert(channelId)

	if err != nil {
		this.Data["json"] = ReturnError(4004, "Request data failed, please try again~")

	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	}
	this.ServeJSON()
}

// Channel page - Get hot play list
// @router /channel/hot [*]
func (this *VideoController) ChannelHotList() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified channel")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelHotList(channelId, 9)

	if err != nil {
		this.Data["json"] = ReturnError(4004, "Not found videos")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	}
	this.ServeJSON()
}

// Channel page- Get channel recommend video by region
// @router /channel/recommend/region [*]
func (this *VideoController) ChannelRecommendRegionList() {
	channelId, _ := this.GetInt("channelId")
	regionId, _ := this.GetInt("regionId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified channel")
		this.ServeJSON()
	}
	if regionId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified region")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "Not found videos")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	}
	this.ServeJSON()
}

// Get video by channel
// @router /channel/recommend/type [*]
func (this *VideoController) GetChannelRecommendTypeList() {
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified channel")
		this.ServeJSON()
	}
	if typeId == 0 {
		this.Data["json"] = ReturnError(4002, "Must specified channel type")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
	}
	this.ServeJSON()
}
