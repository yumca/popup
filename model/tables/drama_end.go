package tables

type DramaEnd struct {
	Id             int    `json:"id"`
	Drama_id       int    `json:"drama_id"`
	End_words      string `json:"end_words"`
	End_factor     int    `json:"end_factor"`
	Showend_factor int    `json:"showend_factor"`
	End_level      int    `json:"end_level"`
	Ctime          int    `json:"ctime"`
}

func (DramaEnd) TableName() string {
	return "drama_end"
}
