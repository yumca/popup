package library

import (
	"encoding/json"
	"errors"
	// "net/http"
	// "github.com/gorilla/websocket"
)

type db struct {
	DbHost   string
	DbUser   string
	DbPwd    string
	DbPort   int
	DbPrefix string
	DbType   string
	Database string
}

type setting struct {
	ServerName string
	LogFile    string
	PidFile    string
}

type Config struct {
	Db      db
	Setting setting
}

func GetConf() (Conf Config, err error) {
	Conf, err = GetConfInfo()
	if err != nil {
		return
	}
	return
}

func GetConfInfo() (conf Config, err error) {
	errJson := json.Unmarshal([]byte(conf_json), &conf)
	if errJson != nil {
		err = errors.New("读取Conf配置错误")
		return
	}
	return
}

var conf_json = `{
    "Db": {
        "DbHost": "127.0.0.1",
        "DbUser": "root",
        "DbPwd": "root",
        "DbPort": 3306,
        "DbPrefix": "",
        "DbType": "sqlite",
        "Database": "test"
    },
    "Setting": {
        "ServerName": "Pupup",
        "LogFile": "./go_pupup.log",
        "PidFile": "./go_pupup.pid"
    }
}
`
