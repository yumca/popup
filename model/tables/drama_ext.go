package tables

type DramaExt struct {
	Id       int    `json:"id"`
	Drama_id int    `json:"drama_id"`
	Images   string `json:"images"`
	Ctime    int    `json:"ctime"`
}

func (DramaExt) TableName() string {
	return "drama_ext"
}
