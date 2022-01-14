package tables

type GameUserAction struct {
	Id       int `json:"id"`
	Game_id  int `json:"game_id"`
	Place_id int `json:"place_id"`
	Uid      int `json:"uid"`
	Tuid     int `json:"tuid"`
	Action   int `json:"action"`
	Ctime    int `json:"ctime"`
}

func (GameUserAction) TableName() string {
	return "game_user_action"
}
