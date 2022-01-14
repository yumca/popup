package tables

type Drama struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Context string `json:"context"`
	Ctime   int    `json:"ctime"`
}

func (Drama) TableName() string {
	return "drama"
}
