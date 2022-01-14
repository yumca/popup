package tables

type GameUserChat struct {
	Id        int    `json:"id"`
	Game_id   int    `json:"game_id"`
	Uid       int    `json:"uid"`
	Chat_uids string `json:"chat_uids"`
	Ctime     int    `json:"ctime"`
}

func (GameUserChat) TableName() string {
	return "game_user_chat"
}
