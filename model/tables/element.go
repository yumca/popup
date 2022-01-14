package tables

type Element struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Width       string `json:"width"`
	Length      string `json:"length"`
	Coordinate  string `json:"coordinate"`
	Quality     string `json:"quality"`
	Stiffness   int    `json:"stiffness"`
	Tenacity    int    `json:"tenacity"`
	Sharp       int    `json:"sharp"`
	Melti_point int    `json:"melti_point"`
	Steam       int    `json:"steam"`
	Electric    int    `json:"electric"`
	Desc        string `json:"desc"`
	Ctime       int    `json:"ctime"`
}

func (Element) TableName() string {
	return "element"
}
