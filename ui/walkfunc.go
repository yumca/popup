package ui

import (
	"popup/model/tables"
	"regexp"
	"time"

	"github.com/lxn/walk"
)

func delete_meeting() {
	var meeting tables.Meeting
	if len(lb.SelectedIndexes()) > 0 {
		// var map_items = make(map[int]contentEntry)
		for _, idx := range lb.SelectedIndexes() {
			meeting.Id = model.items[idx].id
			meeting.Delete()
			// map_items[idx] = model.items[idx]
		}
		var items []contentEntry
		mettings := get_meetings()
		//循环给item列表赋值
		for _, v := range mettings {
			items = append(items, contentEntry{v.Id, time.UnixMilli(int64(v.Timestamp)).Format("2006-01-02 15:04"), v.Content, func(v tables.Meeting) string {
				if v.Notify == 1 {
					return "已通知"
				} else {
					return "未通知"
				}
			}(v)})
		}
		// var tmpitems []contentEntry
		// for i, v := range model.items {
		// 	if _, ok := map_items[i]; !ok {
		// 		tmpitems = append(tmpitems, v)
		// 	}
		// }
		model.items = items
		model.PublishItemsReset()
	}
}

func update_meeting() {
	// var meeting tables.Meeting
	// err := json.Unmarshal([]byte(update_json), &meeting)
	// meeting.Update("notify", meeting.Notify)
}

func get_meetings() []tables.Meeting {
	var meeting tables.Meeting
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	ti := t.UnixMilli()
	return meeting.GetMeetingsByParams("timestamp > ?", ti, "notify asc,timestamp asc,id asc")
}

func save_meeting() {
	text := input.Text()
	reg := regexp.MustCompile(`\d{1,2}(:\d{1,2})`)
	res := reg.FindAllString(text, -1)
	if len(res) > 0 {
		input.SetText("")
		var meeting tables.Meeting
		// var cstSh = time.FixedZone("CST", 8*3600) //东八区
		timeStr := time.Now().UTC().Format("2006-01-02")
		t, _ := time.Parse("2006-01-02 15:04", timeStr+" "+res[0])
		meeting.Timestamp = int(t.UTC().UnixMilli())
		meeting.Content = text
		meeting.Notify = 0
		meeting.Save()
		var items []contentEntry
		mettings := get_meetings()
		//循环给item列表赋值
		for _, v := range mettings {
			items = append(items, contentEntry{v.Id, time.UnixMilli(int64(v.Timestamp)).Format("2006-01-02 15:04"), v.Content, func(v tables.Meeting) string {
				if v.Notify == 1 {
					return "已通知"
				} else {
					return "未通知"
				}
			}(v)})
		}
		model.items = items
		model.PublishItemsReset()
	} else {
		walk.MsgBox(UiMainWindow, "匹配错误", "未匹配到正确时间格式|小时:分钟", walk.MsgBoxIconInformation)
	}
}
