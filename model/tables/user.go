package tables

import "popup/model"

type User struct {
	Id       int    `json:"id"`
	Uname    string `json:"uname"`
	Wsfd     int    `json:"wsfd"`
	Tcpfd    int    `json:"tcpfd"`
	Loginkey string `json:"loginkey"`
	Ctime    int    `json:"ctime"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Create() {
	db := model.GetDb()
	db.Create(u)
}

func (u *User) Save() {
	db := model.GetDb()
	db.Save(u)
}

func (u *User) GetUserByKey(loginkey string) *User {
	if loginkey == "" {
		return u
	}
	db := model.GetDb()
	db.Where("loginkey = ?", loginkey).First(u)
	return u
}

func (u User) GetUsers(params ...interface{}) []User {
	var users []User
	db := model.GetDb()
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.Find(&users)
	return users
}

func (u *User) GetUserInfo(params ...interface{}) *User {
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

func (u *User) Update(column string, value interface{}, params ...interface{}) *User {
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

func (u *User) Updates(datas ...map[string]interface{}) *User {
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
