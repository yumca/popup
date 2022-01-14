package ui

import (
	"popup/view"
	"time"

	"github.com/go-vgo/robotgo"
)

func jsbind() {
	//监控事件
	event_monitor()
	// Bind Go functions to JS
	Lorcaui.Bind("lorca_loadcontent", lorca_loadcontent)
	Lorcaui.Bind("lorca_loaduri", lorca_loaduri)
	Lorcaui.Bind("lorca_localjs", lorca_localjs)
	Lorcaui.Bind("lorca_notification", lorca_notification)
}

func lorca_notification(t string) {
	robotgo.ShowAlert("弹窗", "消息", "按钮1", "退出按钮")
	// timeLayout := time.Now().Format("2006-01-02")
	// timeLayout = timeLayout + " " + t + ":00"
	// loc, _ := time.LoadLocation("Local") //获取时区
	// tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", timeLayout, loc)
	// timestamp := tmp.Unix()
	// Lorcaui.Eval(view.GetJs(title))
	// return view.GetJs(title)
}

func lorca_localjs(title string) string {
	if view.GetJs(title) != "" {
		Lorcaui.Eval(view.GetJs(title))
	}
	return view.GetJs(title)
}

func lorca_loadcontent(title string) {
	Lorcaui.Reload()
	frameTree, _ := Lorcaui.GetFrameTree()
	Lorcaui.SetContent(frameTree.FrameTree.Frame.Id, view.GetView(title))
}

func lorca_loaduri(url string) {
	Lorcaui.Load(url)
}

//事件监控
func event_monitor() {
	//注册导航跳转事件
	Lorcaui.SetEvent("Page.frameNavigated")
	//执行禁止操作
	runBan()
	go func() {
		for {
			//监控导航跳转事件
			unixMicro, _ := Lorcaui.PopEvent("Page.frameNavigated")
			//如果有导航跳转事件冒出  执行禁止操作
			if unixMicro > 0 {
				runBan()
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func runBan() {
	Lorcaui.Eval(`document.onkeydown = function(){
			//禁止F12
			if(window.event && (window.event.keyCode === 123 || window.event.which === 123)) {
				window.event.cancelBubble = true;
	    		window.event.returnValue = false;
	    		window.event.keyCode = 0;
			}
			if(window.event && window.event.keyCode == 13) {
				window.event.keyCode = 505;
			}
			if(window.event && window.event.ctrlKey && window.event.shiftKey && window.event.keyCode==73) {
				window.event.cancelBubble = true;
	    		window.event.returnValue = false;
	    		window.event.keyCode = 0;
				e.preventDefault()
			}
			//屏蔽F11
			if (window.event.keyCode == 122) {
				window.event.cancelBubble = true;
				window.event.keyCode = 0;
				window.event.returnValue = false;
			}
			//屏蔽 Ctrl+n
			if (window.event.ctrlKey && window.event.keyCode == 78) {
				window.event.returnValue = false;
			}
			//屏蔽 shift+F10
			if (window.event.shiftKey && window.event.keyCode == 121) {
				window.event.returnValue = false;
			}
			//屏蔽 shift 加鼠标左键新开一网页
			if (window.event.srcElement.tagName == "A" && window.event.shiftKey){
				window.event.returnValue = false;
			}
		}`)
	// 禁止右键菜单
	Lorcaui.Eval(`document.oncontextmenu = function (event){
			if(window.event){
				event = window.event;
			}
			try{
				var the = event.srcElement;
				if (!((the.tagName == "INPUT" && the.type.toLowerCase() == "text") || the.tagName == "TEXTAREA")){
					return false;
				}
				return true;
			}catch (e){
				return false;
			}
		}`)
	//屏蔽F1帮助
	Lorcaui.Eval(`window.onhelp = function (){
			return false;
	}`)
}
