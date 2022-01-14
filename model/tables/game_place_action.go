package tables

type GamePlaceAction struct {
	Id          int    `json:"id"`
	Game_id     int    `json:"game_id"`
	Place_id    int    `json:"place_id"`
	Uid         int    `json:"uid"`
	To_users    string `json:"to_users"`
	Action_type int    `json:"action_type"`
	Content     string `json:"content"`
	Ctime       int    `json:"ctime"`
}

func (GamePlaceAction) TableName() string {
	return "game_place_action"
}
