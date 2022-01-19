package view

func GetView(title string) string {
	var html string
	switch title {
	case "main":
		html = viewMain()
	case "test":
		html = viewTest()
	default:
		html = "data:text/html,<html><head><title>googleapp</title></head></html>"
	}
	return html
}

func GetJs(title string) string {
	var js string
	switch title {
	case "jquery":
		js = jquery()
	case "jquery-cookie":
		js = jquery_cookie()
	default:
		js = ""
	}
	return js
}

func viewMain() string {
	return `<html>

	<head>
		<title>main</title>
		<meta charset="UTF-8">
	</head>

	<body>
		<!--<button onclick="lorca_loadcontent('main')">刷新</button>-->
		<!--<button onclick="lorca_alert('弹窗','弹窗内容')">test</button>-->
		<div><textarea id="textarea" style="height: 40%; width: 100%;"></textarea></div>
		<div id="list">
			<div>无</div>
		</div>
	</body>
	<script>
		var notify_list = []
		lorca_localjs("jquery").then(function (res) {
			eval(res)
			run()
			setInterval(notify, 10000)
		})
		function run() {
			lorca_get_meetings().then(function (res) {
				if (res != "") {
					notify_list = JSON.parse(res)
				} else {
					notify_list = []
				}
				changelist()
			})
			// 14:15，2号会议室，项目开发子系统和产品项目子系统对接
			$('#textarea').blur(function () {
				var reg = /\d{1,2}(:\d{1,2})/;
				var reg_res = reg.exec($('#textarea').val())
				if (reg_res != null) {
					var textarea = $('#textarea').val()
					$('#textarea').val("")
					var myDate = new Date;
					var year = myDate.getFullYear(); //获取当前年
					var mon = myDate.getMonth() + 1; //获取当前月
					var date = myDate.getDate(); //获取当前日
					var date_str = year + "-" + mon + "-" + date + " " + reg_res[0] + ":00"
					var lt = year + "-" + mon + "-" + date + "09:00:00"
					var notifytime = (new Date(date_str));//把当前日期变成毫秒时间戳
					if (notifytime < lt) {
						notifytime += 43200 * 1000
					}
					var val = {
						content: textarea,
						timestamp: notifytime - 120000,
						notify: 0
					}
					lorca_save_meeting(JSON.stringify(val)).then(function (err) {
						lorca_get_meetings().then(function (res) {
							if (res != "") {
								notify_list = JSON.parse(res)
							} else {
								notify_list = []
							}
							changelist()
						})
					})
				}
			});
		}

		function changelist() {
			if (notify_list.length > 0) {
				var html = "";
				for (var i = 0; i < notify_list.length; i++) {
					html += '<div>' + notify_list[i].content + '<button onclick="delete_meeting(' + notify_list[i].id + ')">删除</button>'
					if (notify_list[i].notify == true) {
						html += '<span style="color: red;">已通知</span>'
					} else {
						html += '<span style="color: green;">未通知</span>'
					}
					var myDate = new Date(notify_list[i].timestamp);
					var year = myDate.getFullYear(); //获取当前年
					var mon = myDate.getMonth() + 1; //获取当前月
					var date = myDate.getDate(); //获取当前日
					var hour = myDate.getHours(); //获取当前小时
					var minute = myDate.getMinutes(); //获取当前分钟
					var date_str = year + "-" + mon + "-" + date + " " + hour + ":" + minute
					html += '<span>通知时间：' + date_str + '</span>'
					html += '</div>'
				}
				$('#list').html(html)
			}
		}

		function delete_meeting(id) {
			lorca_delete_meeting(id).then(function (res) {
				lorca_get_meetings().then(function (res) {
					if (res != "") {
						notify_list = JSON.parse(res)
					} else {
						notify_list = []
					}
					changelist()
				})
			})
		}

		function notify(type, title, message) {
			if (notify_list.length > 0) {
				for (var i = 0; i < notify_list.length; i++) {
					var now = (new Date()).getTime()
					if (notify_list[i].notify == 0 && notify_list[i].timestamp <= now) {
						notify_list[i].notify = 1
						lorca_alert("开会通知", notify_list[i].content)
						lorca_update_meeting(JSON.stringify(notify_list[i])).then(function (res) {
							if (res == "") {
								lorca_get_meetings().then(function (res) {
									if (res != "") {
										notify_list = JSON.parse(res)
									} else {
										notify_list = []
									}
									changelist()
								})
							}
						})

					}
				}
			}
		}
	</script>

	</html>

	`
}

func viewMain_bak() string {
	return `<html>
	<head>
		<title>main</title>
		<meta charset="UTF-8">
	</head>
	<body>
		<div>
		<button onclick="lorca_loaduri('http://172.16.3.23:8080')">jumpTo：172.16.3.23:8080</button>
		<button id="a" onclick="localjs()">测试js是否加载</button>
		<!--<button onclick="lorca_loadcontent('test')">加载test页面</button>-->
		<button onclick="lorca_alert('弹窗','test')">弹窗</button>
	</body>
	<script>
	lorca_localjs("jquery")
	</script>
	</html>
	`
}
