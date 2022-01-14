package tables

type GamePlaceLog struct {
	Id       int    `json:"id"`
	Game_id  int    `json:"game_id"`
	Place_id int    `json:"place_id"`
	Uid      int    `json:"uid"`
	Content  string `json:"content"`
	Ctime    int    `json:"ctime"`
}

func (GamePlaceLog) TableName() string {
	return "game_place_log"
}
