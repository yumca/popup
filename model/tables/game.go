package tables

import "popup/model"

type Game struct {
	Id       int    `json:"id"`
	Drama_id int    `json:"drama_id"`
	Uid      int    `json:"uid"`
	Title    string `json:"title"`
	Play     int    `json:"play"`
	Ctime    int    `json:"ctime"`
}

func (Game) TableName() string {
	return "game"
}

func (g Game) GetGameList(raw string, param ...interface{}) []Game {
	var res []Game
	db := model.GetDb()
	if raw != "" {
		db = db.Where(raw, param...)
	}
	db.Find(&res)
	return res
}

func (g *Game) GetGameInfo(params ...interface{}) *Game {
	db := model.GetDb()
	if len(params) > 0 && params[0] != "" {
		if len(params) > 1 {
			db = db.Where(params[0], params[1:]...)
		} else {
			db = db.Where(params[0])
		}
	}
	db.First(g)
	return g
}

func (g *Game) Create() {
	db := model.GetDb()
	db.Create(g)
}

func (g *Game) Save() {
	db := model.GetDb()
	db.Save(g)
}
