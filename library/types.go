package library

type ResultData struct {
	Status    int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}
