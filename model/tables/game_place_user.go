package tables

type GamePlaceUser struct {
	Id       int `json:"id"`
	Game_id  int `json:"game_id"`
	Place_id int `json:"place_id"`
	Uid      int `json:"uid"`
	Ctime    int `json:"ctime"`
}

func (GamePlaceUser) TableName() string {
	return "game_place_user"
}
