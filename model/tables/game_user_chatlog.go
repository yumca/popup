package tables

type GameUserChatlog struct {
	Id      int    `json:"id"`
	Game_id int    `json:"game_id"`
	Uid     int    `json:"uid"`
	Tuid    int    `json:"tuid"`
	Content string `json:"content"`
	Ctime   int    `json:"ctime"`
}

func (GameUserChatlog) TableName() string {
	return "game_user_chatlog"
}
