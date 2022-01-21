package ui

import (
	"encoding/json"
	"popup/model/tables"
	"time"
)

func delete_meeting() {
	var meeting tables.Meeting
	if len(lb.SelectedIndexes()) > 0 {
		var map_items = make(map[int]contentEntry)
		for _, idx := range lb.SelectedIndexes() {
			meeting.Id = model.items[idx-1].id
			meeting.Delete()
			map_items[idx-1] = model.items[idx-1]
		}
		var tmpitems []contentEntry
		for i, v := range model.items {
			if _, ok := map_items[i]; !ok {
				tmpitems = append(tmpitems, v)
			}
		}
		model.items = tmpitems
		model.PublishItemsReset()
	}
}

func update_meeting(update_json string) string {
	var meeting tables.Meeting
	err := json.Unmarshal([]byte(update_json), &meeting)
	if err != nil {
		return err.Error()
	}
	meeting.Update("notify", meeting.Notify)
	return ""
}

func get_meetings() []tables.Meeting {
	var meeting tables.Meeting
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	ti := t.UnixMilli()
	return meeting.GetMeetingsByParams("timestamp > ?", ti, "notify asc,timestamp asc,id asc")
}

func save_meeting() {
	text = input.Text()
	var meeting tables.Meeting
	err := json.Unmarshal([]byte(save_json), &meeting)
	if err != nil {
		return err.Error()
	}
	meeting.Save()
}
