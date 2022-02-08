package ui

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"popup/model/tables"
	"regexp"
	"time"

	"github.com/lxn/walk"
)

// 14:15，2号会议室，项目开发子系统和产品项目子系统对接

var ticker *time.Ticker
var cancelticker chan bool

func delete_meeting() {
	var meeting tables.Meeting
	if len(lb.SelectedIndexes()) > 0 {
		for _, idx := range lb.SelectedIndexes() {
			meeting.Id = model.items[idx].id
			meeting.Delete()
		}
		reflash()
	}
}

func update_meeting(c contentEntry) {
	var meeting tables.Meeting
	meeting.Id = c.id
	meeting.Update("notify", 1)
}

func get_meetings() []tables.Meeting {
	var meeting tables.Meeting
	timeStr := time.Now().Format("2006-01-02")
	// L, _ := time.LoadLocation("Asia/Shanghai")
	L := time.FixedZone("CST", 8*3600)
	t, _ := time.ParseInLocation("2006-01-02", timeStr, L)
	return meeting.GetMeetingsByParams("timestamp > ?", t.UnixMilli(), "notify asc,timestamp asc,id asc")
}

func save_meeting() {
	text := input.Text()
	reg := regexp.MustCompile(`\d{1,2}(:\d{1,2})`)
	res := reg.FindAllString(text, -1)
	if len(res) > 0 {
		input.SetText("")
		var meeting tables.Meeting
		timeStr := time.Now().Format("2006-01-02")
		// L, _ := time.LoadLocation("Asia/Shanghai")
		L := time.FixedZone("CST", 8*3600)
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" "+res[0]+":00", L)
		worktime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 09:00:00", L)
		if t.UnixMilli() < worktime.UnixMilli() {
			meeting.Timestamp = int(t.UnixMilli()) + 43200000 - 120000
		} else {
			meeting.Timestamp = int(t.UnixMilli()) - 120000
		}
		meeting.Content = text
		meeting.Notify = 0
		meeting.Save()
		reflash()
	} else {
		walk.MsgBox(UiMainWindow, "匹配错误", "未匹配到正确的时间格式 | 小时:分钟", walk.MsgBoxIconInformation)
	}
}

func notifyTicker() {
	ticker = time.NewTicker(time.Second * 10)
	cancelticker = make(chan bool, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				nowTime := time.Now().UnixMilli()
				for _, v := range model.items {
					// L, _ := time.LoadLocation("Asia/Shanghai")
					L := time.FixedZone("CST", 8*3600)
					t, _ := time.ParseInLocation("2006-01-02 15:04:05", v.timestamp, L)
					if nowTime >= t.UnixMilli() && v.notify == "未通知" {
						err := doNotification("windows通知", v.content)
						if err != nil {
							walk.MsgBox(UiMainWindow, "通知错误", "通知失败："+err.Error(), walk.MsgBoxIconInformation)
							break
						}
						update_meeting(v)
						reflash()
					}
				}
				// UiMainWindow.Synchronize(func() {
				// 	trackLatest := lb.ItemVisible(len(model.items)-1) && len(lb.SelectedIndexes()) <= 1
				// 	model.items = append(model.items, contentEntry{1, "1", "Some new stuff.", "sss"})
				// 	index := len(model.items) - 1
				// 	model.PublishItemsInserted(index, index)

				// 	if trackLatest {
				// 		lb.EnsureItemVisible(len(model.items) - 1)
				// 	}
				// })

			case <-cancelticker:
				ticker.Stop()
				break
			}
		}
	}()
}

func doNotification(t, m string) (err error) {
	if err = notifyIcon.ShowCustom(t, m, icon); err != nil {
		return
	}

	return nil
}

func reflash() {
	var items []contentEntry
	mettings := get_meetings()
	//循环给item列表赋值
	for _, v := range mettings {
		items = append(items, contentEntry{v.Id, time.UnixMilli(int64(v.Timestamp)).Format("2006-01-02 15:04:05"), v.Content, func(v tables.Meeting) string {
			if v.Notify == 1 {
				return "已通知"
			} else {
				return "未通知"
			}
		}(v)})
	}
	model.items = items
	model.PublishItemsReset()
}

type open struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Opentime int    `json:"opentime"`
}

func getOpen() {
	url := "http://mbstr.hundian.club/dingding/api/get_open"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var st open
	err = json.Unmarshal(data, &st) //第二个参数要地址传递
	if err != nil {
		return
	}
	if st.Code == 0 {
		var meeting tables.Meeting
		L := time.FixedZone("CST", 8*3600)
		t1, _ := time.ParseInLocation("2006-01-02", time.UnixMilli(int64(st.Opentime*1000)).Format("2006-01-02"), L)
		t2 := int(t1.UnixMilli()) + 86400000 - 1000
		meeting.GetMeetingInfo("content LIKE ? AND timestamp > ? AND timestamp < ?", "今天开机%", t1.UnixMilli(), t2)
		if meeting.Id < 1 {
			meeting.Timestamp = (st.Opentime + 30600) * 1000
			meeting.Content = st.Msg
			meeting.Notify = 0
			meeting.Save()
			reflash()
		}
	} else if st.Code == 1 {
		walk.MsgBox(UiMainWindow, "获取开机时间失败", st.Msg, walk.MsgBoxIconInformation)
	}
}
