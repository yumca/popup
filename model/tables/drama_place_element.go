package tables

type DramaPlaceElement struct {
	Id           int `json:"id"`
	Drama_id     int `json:"drama_id"`
	Place_id     int `json:"place_id"`
	Element_id   int `json:"element_id"`
	X            int `json:"x"`
	Y            int `json:"y"`
	Keyword      int `json:"keyword"`
	Keylevel     int `json:"keylevel"`
	Weapon       int `json:"weapon"`
	Weapon_level int `json:"weapon_level"`
	Ctime        int `json:"ctime"`
}

func (DramaPlaceElement) TableName() string {
	return "drama_place_element"
}
