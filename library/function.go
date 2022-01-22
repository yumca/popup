package library

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/thinkeridea/go-extend/exnet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

/**
 * 去空
 */
func TrimEmpty(a []string) (ret []string) {
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

/**
 * 日志
 */
func DoMyLogs(msg string, types string, mode string, data string) {
	//MyLogs.NewMyLogs()
	//$trace = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS, 1)
	//$this- > logs- > doLog($msg, $data, $type, current($trace))
}

/**
 * 获取ip
 */
func Getip(r *http.Request) string {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	if ip == "" {
		ip = "0.0.0.0"
	}
	return ip
}

func GetClientOs(r *http.Request) string {
	os := "other"
	userAgent := strings.ToLower(r.Header.Get("User-Agent"))
	if re := strings.IndexAny(userAgent, "iphone"); re != -1 {
		os = "iphone"
	} else if re := strings.IndexAny(userAgent, "android"); re != -1 {
		os = "android"
	} else if re := strings.IndexAny(userAgent, "micromessenger"); re != -1 {
		os = "weixin"
	} else if re := strings.IndexAny(userAgent, "ipad"); re != -1 {
		os = "ipad"
	} else if re := strings.IndexAny(userAgent, "ipod"); re != -1 {
		os = "ipod"
	} else if re := strings.IndexAny(userAgent, "windows nt"); re != -1 {
		os = "pc"
	}
	return os
}

//func GetItemId() string {
//    $hour = date("z") * 24 + date("H");
//    $hour = str_repeat("0", 4 - strlen($hour)) . $hour;
//    //	echo date("y") . $hour . PHP_EOL;
//    return date("y") . $hour . getRandNumber(10);
//}

func GetItemId() string {
	return ""
}

//返回毫秒时间戳 10+3
func GetMillisecond() string {
	// rand.Seed(time.Now().UnixNano())
	return strconv.FormatInt(time.Now().UnixMilli(), 10) //+ strconv.Itoa(rand.Intn(100))
}

//返回时间戳 10
func GetSecond() int {
	// rand.Seed(time.Now().UnixNano())
	return int(time.Now().Unix()) //+ strconv.Itoa(rand.Intn(100))
}

//设置api返回数据格式
func ApiResult(status int, msg string, data interface{}) ResultData {
	return ResultData{
		status,
		msg,
		data,
		strconv.FormatInt(time.Now().Unix(), 10),
	}
}

//生成32位md5字串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GBK 转 UTF-8
func GbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), unicode.UTF8.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return s
	}
	return d
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		// log.Fatal(e)
		return s
	}
	return d
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home + "\\"
	}
	return os.Getenv("HOME") + "/"
}

func UserLocalLowDir() string {
	if runtime.GOOS == "windows" {
		path := UserHomeDir() + "AppData\\LocalLow\\"
		return path
	}
	return UserHomeDir() + "LocalLow/"
}

func UserLocalDir() string {
	if runtime.GOOS == "windows" {
		path := UserHomeDir() + "AppData\\Local\\"
		return path
	}
	return UserHomeDir() + "Local/"
}

func ProgramDir() string {
	conf, err := GetConf()
	if err != nil {
		return ""
	}
	path := ""
	if runtime.GOOS == "windows" {
		path = UserLocalLowDir() + conf.Setting.ServerName + "\\"
	} else {
		path = UserLocalLowDir() + conf.Setting.ServerName + "/"
	}
	if c := MkdirAll(path); !c {
		return ""
	}
	return path
}

func MkdirAll(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		//递归创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

//
///**
// * 生成固定长度的随机数
// *
// * @param int $length
// * @return string
// */
//function getRandNumber($length = 6) {
//    $num = "";
//    if ($length >= 10) {
//        $t = intval($length / 9);
//        $tail = $length % 9;
//        for ($i = 1; $i <= $t; $i ++) {
//            $num .= substr(mt_rand("1" . str_repeat("0", 9), str_repeat("9", 10)), 1);
//        }
//        $num .= substr(mt_rand("1" . str_repeat("0", $tail), str_repeat("9", $tail + 1)), 1);
//        return $num;
//    } else {
//        return substr(mt_rand("1" . str_repeat("0", $length), str_repeat("9", $length + 1)), 1);
//    }
//}
//
///**
// * ws返回格式化
// */
//function ws_return($signal, $code = 1, $msg = "succ", $data = []) {
//    return [
//        "signal" => $signal,
//        "code" => intval($code),
//        "msg" => $msg,
//        "serverTime" => time(),
//        "data" => $data
//    ];
//}
