package tables

import "popup/model"

type GameUser struct {
	Id      int `json:"id"`
	Game_id int `json:"game_id"`
	Uid     int `json:"uid"`
	Role_id int `json:"role_id"`
	Killer  int `json:"killer"`
	Ctime   int `json:"ctime"`
}

func (GameUser) TableName() string {
	return "game_user"
}

func (gu GameUser) GetGameUsers(params ...interface{}) []GameUser {
	var res []GameUser
	db := model.GetDb()
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.Find(&res)
	return res
}

func (gu *GameUser) GetGameUserInfo(params ...interface{}) *GameUser {
	db := model.GetDb()
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.First(gu)
	return gu
}

func (gu *GameUser) Create() *GameUser {
	db := model.GetDb()
	db.Create(gu)
	return gu
}

func (gu *GameUser) Delete(params ...interface{}) *GameUser {
	db := model.GetDb()
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.Delete(gu)
	return gu
}
