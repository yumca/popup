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
		<div>
		<button onclick="lorca_loaduri('http://172.16.3.23:8080')">jumpTo：172.16.3.23:8080</button>
		<button id="a" onclick="localjs()">测试js是否加载</button>
		<!--<button onclick="lorca_loadcontent('test')">加载test页面</button>-->
		<button onclick="lorca_notification('test')">弹窗</button>
	</body>
	<script>
	lorca_localjs("jquery")
	function localjs(){
			if($("#a").length > 0){
				alert(1)
			}
	}
	</script>
	</html>
	`
}
