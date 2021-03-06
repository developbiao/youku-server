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

// Get videos by request parameters
// @router /channel/video [*]
func (this *VideoController) ChannelVideo() {
	// Get channel ID
	channelId, _ := this.GetInt("channelId")
	// Get region ID
	regionId, _ := this.GetInt("regionId")
	// Get type ID
	typeId, _ := this.GetInt("typeId")
	// Get state
	end := this.GetString("end")
	// Get sort
	sort := this.GetString("sort")
	// Get paginate
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "Must specified channel")
		this.ServeJSON()
	}

	// Default limit 12
	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
	}
	this.ServeJSON()

}

// Get Video detail
// @router /video/info [*]
func (this *VideoController) VideoInfo() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频ID")
		this.ServeJSON()
	}
	video, err := models.GetVideoInfo(videoId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", video, 1)
	}
	this.ServeJSON()
}

// Get video episodes list
// @router /video/episodes/list [*]
func (this *VideoController) VideoEpisodesList() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频ID")
		this.ServeJSON()
	}

	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", episodes, num)
	}
	this.ServeJSON()
}

// My video manage
// @router /user/video [*]
func (this *VideoController) UserVideo() {
	uid, _ := this.GetInt("uid")
	if uid == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定用户")
		this.ServeJSON()
	}

	num, videos, err := models.GetUserVideo(uid)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
	}
	this.ServeJSON()
}

// Save user upload video information
// @router /video/save [*]
func (this *VideoController) VideoSave() {
	playUrl := this.GetString("playUrl")
	title := this.GetString("title")
	subTitle := this.GetString("subTitle")
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")
	regionId, _ := this.GetInt("regionId")
	uid, _ := this.GetInt("uid")
	aliyunVideoId := this.GetString("aliyunVideoId")
	if uid == 0 {
		this.Data["json"] = ReturnError(4001, "请先登录")
		this.ServeJSON()
	}
	if playUrl == "" {
		this.Data["json"] = ReturnError(4002, "视频地址不能为空")
		this.ServeJSON()
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid, aliyunVideoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(5000, err)
		this.ServeJSON()
	}

}
