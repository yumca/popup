package tables

import (
	"popup/model"
	"strings"
)

type Meeting struct {
	Id        int    `json:"id"`
	Timestamp int    `json:"timestamp"`
	Content   string `json:"content"`
	Notify    int    `json:"notify"`
}

func (Meeting) TableName() string {
	return "meeting"
}

func (u *Meeting) Create() {
	db := model.GetDb()
	db.Create(u)
}

func (u *Meeting) Save() {
	db := model.GetDb()
	db.Save(u)
}

func (u *Meeting) Delete() {
	db := model.GetDb()
	db.Delete(u)
}

func (u Meeting) GetMeetings(params ...interface{}) []Meeting {
	var Meetings []Meeting
	db := model.GetDb()
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.Find(&Meetings)
	return Meetings
}

func (u Meeting) GetMeetingsByParams(params ...interface{}) []Meeting {
	var Meetings []Meeting
	db := model.GetDb()
	if len(params) > 0 {
		if params[0] != "" {
			c := strings.Count(params[0].(string), "?")
			if c > 0 {
				db = db.Where(params[0], params[1:c+1]...)
			} else {
				db = db.Where(params[0])
			}
			params = params[c+1:]
		} else {
			params = params[1:]
		}
		if len(params) > 0 && params[0] != "" {
			db = db.Order(params[0])
		}
	}
	db.Find(&Meetings)
	return Meetings
}

func (u *Meeting) GetMeetingInfo(params ...interface{}) *Meeting {
	db := model.GetDb()
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.First(u)
	return u
}

func (u *Meeting) Update(column string, value interface{}, params ...interface{}) *Meeting {
	db := model.GetDb().Model(u)
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.Update(column, value)
	return u
}

func (u *Meeting) Updates(datas ...map[string]interface{}) *Meeting {
	db := model.GetDb()
	update := datas[0]
	if len(datas) > 1 {
		update = datas[1]
		db.Model(u).Where(datas[0]).Updates(update)
	} else {
		db.Model(u).Updates(update)
	}
	return u
}
