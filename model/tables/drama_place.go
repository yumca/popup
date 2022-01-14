package tables

type DramaPlace struct {
	Id         int    `json:"id"`
	Drama_id   int    `json:"drama_id"`
	Place_name string `json:"place_name"`
	Width      int    `json:"width"`
	Length     int    `json:"length"`
	Height     int    `json:"height"`
	Coordinate string `json:"coordinate"`
	Desc       string `json:"desc"`
	Ctime      int    `json:"ctime"`
}

func (DramaPlace) TableName() string {
	return "drama_place"
}
