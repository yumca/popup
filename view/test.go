package view

func viewTest() string {
	return `<html>
	<head>
		<title>test</title>
		<meta charset="UTF-8">
	</head>
	<body>
		<button onclick="lorca_loaduri('http://dev.oa3.com/')">jumpTo：dev.oa3.com</button>
		<button onclick="lorca_loaduri('http://172.16.3.23:8080')">jumpTo：172.16.3.23:8080</button>
		<button id="a" onclick="localjs()">测试js是否加载</button>
	</body>
	<script>
	function localjs(){
			if($("#a").length > 0){
				alert($("#a").length,$("#a").innderHTML)
			}
	}
	</script>
	</html>
	`
}
