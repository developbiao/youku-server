package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Barrage struct {
	Id          int
	Content     string
	CurrentTime int
	AddTime     int64
	UserId      int
	Status      int
	EpisodesId  int
	VideoId     int
}

type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

func init() {
	orm.RegisterModel(new(Barrage))
}

func BarrageList(episodesId int, startTime int, endTime int) (int64, []BarrageData, error) {
	o := orm.NewOrm()
	var barrages []BarrageData
	num, err := o.Raw("SELECT `id`, `content`, `current_time` FROM `barrage`"+
		" WHERE `status` = 1 AND episodes_id=? AND `current_time`>=? AND `current_time`< ? ORDER BY `current_time` ASC",
		episodesId, startTime, endTime).
		QueryRows(&barrages)
	if err != nil {
		fmt.Println(err)
	}
	if num == 0 {
		fmt.Println("Not found data")
	} else {
		fmt.Println("Got data!")
		fmt.Println(barrages)
	}
	return num, barrages, err
}

func SaveBarrage(episodesId int, videoId int, currentTime int, userId int, content string) error {
	o := orm.NewOrm()
	var barrage Barrage
	barrage.EpisodesId = episodesId
	barrage.VideoId = videoId
	barrage.CurrentTime = currentTime
	barrage.UserId = userId
	barrage.Content = content
	barrage.AddTime = time.Now().Unix()
	barrage.Status = 1
	_, err := o.Insert(&barrage)
	return err
}
