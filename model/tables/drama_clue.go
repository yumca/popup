package tables

type DramaClue struct {
	Id             int    `json:"id"`
	Drama_id       int    `json:"drama_id"`
	Word           string `json:"word"`
	Keyword        string `json:"keyword"`
	Keyword_factor int    `json:"keyword_factor"`
	End_id         int    `json:"end_id"`
	Desc           string `json:"desc"`
	Ctime          int    `json:"ctime"`
}

func (DramaClue) TableName() string {
	return "drama_clue"
}
