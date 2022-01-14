package tables

type DramaRole struct {
	Id        int    `json:"id"`
	Drama_id  int    `json:"drama_id"`
	Role_name string `json:"role_name"`
	Desc      string `json:"desc"`
	Ctime     int    `json:"ctime"`
}

func (DramaRole) TableName() string {
	return "drama_role"
}
