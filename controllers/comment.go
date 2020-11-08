package controllers

import (
	"fmt"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
	EpisodesId   int             `json:"episodesId"`
}

// Get comment list
// @router /comment/list [*]
func (this *CommentController) List() {
	// Get episodes id
	episodesId, _ := this.GetInt("episodesId")
	// Get paginate
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if episodesId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频剧集")
		this.ServeJSON()
	}

	if limit == 0 {
		limit = 12
	}

	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有相关内容~")
		this.ServeJSON()
	} else {
		var data []CommentInfo
		var commentInfo CommentInfo
		// Get uid channel
		uidChan := make(chan int, 12)
		closeChan := make(chan bool, 5)
		resChan := make(chan models.UserInfo, 12)

		//  put the obtained uid in the channel
		go func() {
			for _, v := range comments {
				uidChan <- v.UserId
			}
			close(uidChan)
		}()

		// Process uidChannel inside information
		for i := 0; i < 5; i++ {
			go chanGetUserInfo(uidChan, resChan, closeChan)
		}

		// Judge executed finished, information aggregation
		go func() {
			for i := 0; i < 5; i++ {
				<-closeChan
			}
			close(resChan)
			close(closeChan)

		}()

		userInfoMap := make(map[int]models.UserInfo)
		for r := range resChan {
			userInfoMap[r.Id] = r
		}

		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DateFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			commentInfo.EpisodesId = v.EpisodesId
			// Get user information
			commentInfo.UserInfo, _ = userInfoMap[v.UserId]
			data = append(data, commentInfo)
		}

		this.Data["json"] = ReturnSuccess(0, "success", data, num)
		this.ServeJSON()
	}
}

func chanGetUserInfo(uidChan chan int, resChan chan models.UserInfo, closeChan chan bool) {
	for uid := range uidChan {
		res, err := models.GetUserInfo(uid)
		fmt.Println(res)
		if err == nil {
			resChan <- res
		}
	}
	closeChan <- true
}
