package library

import (
	"encoding/json"
	"errors"

	// "net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

type rd struct {
	Stat        string
	RedisHost   string
	RedisPort   string
	RedisPrefix string
	RedisPwd    string
	RedisDb     int
}

type setting struct {
	ServerName string
	LogFile    string
	PidFile    string
	Daemonize  int
}

type Config struct {
	Db       db
	Redis    rd
	Setting  setting
	ConfPath string
	Server   map[string]string
}

//获取当前文件执行路径
func GetExecPath() string {
	execFile, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(execFile)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

func GetConf() (Conf Config, err error) {
	path := GetExecPath()
	Conf, err = GetConfInfo(path)
	if err != nil {
		return
	}
	Conf.ConfPath = path
	return
}

func GetConfInfo(path string) (conf Config, err error) {
	var tmpConf Config
	errJson := json.Unmarshal([]byte(conf_json), &tmpConf)
	if errJson != nil {
		err = errors.New("读取Conf配置错误")
		return
	}
	conf = tmpConf
	return
}

var conf_json = `{
    "Db": {
        "DbHost": "127.0.0.1",
        "DbUser": "root",
        "DbPwd": "root",
        "DbPort": 3306,
        "DbPrefix": "",
        "DbType": "mysql",
        "Database": "test"
    },
    "Redis": {
        "Stat": "off",
        "RedisHost": "127.0.0.1",
        "RedisPort": "6379",
        "RedisPrefix": "task_",
        "RedisPwd": "",
        "RedisDb": 9
    },
    "Setting": {
        "ServerName": "Pupup",
        "LogFile": "./go_pupup.log",
        "PidFile": "./go_pupup.pid",
        "Daemonize": 0
    },
    "Server": {
        "http": "172.16.4.66:8088",
        "ws": "172.16.4.66:8090",
        "tcp": "172.16.4.66:8070"
    }
}
`
