package tables

type GameLog struct {
	Id       int    `json:"id"`
	Game_id  int    `json:"game_id"`
	Place_id int    `json:"place_id"`
	Users_id string `json:"users_id"`
	Content  string `json:"content"`
	Ctime    int    `json:"ctime"`
}

func (GameLog) TableName() string {
	return "game_log"
}
