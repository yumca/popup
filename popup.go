package main

import "popup/ui"

func main() {
	ui.LorcaNew(480, 320)
	defer ui.LorcaClose()
	<-ui.LorcaDone()

	// ui, err = lorca.New("", "", 480, 320)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer ui.Close()
	// start_bind()
	// // Bind Go functions to JS
	// ui.Bind("toggle", toggle)
	// ui.Bind("reset", reset)
	// // Load HTML after Go functions are bound to JS
	// ui.Load("data:text/html," + url.PathEscape(`
	// <html>
	// 	<body>
	// 		<!-- toggle() and reset() are Go functions wrapped into JS -->
	// 		<div class="timer" onclick="toggle()"></div>
	// 		<button onclick="reset()">Reset</button>
	// 	</body>
	// </html>
	// `))

	// // Start ticker goroutine
	// go func() {
	// 	t := time.NewTicker(100 * time.Millisecond)
	// 	for {
	// 		select {
	// 		case <-t.C: // Every 100ms increate number of ticks and update UI
	// 			ui.Eval(fmt.Sprintf(`document.querySelector('.timer').innerText = 0.1*%d`,
	// 				atomic.AddUint32(&ticks, 1)))
	// 		case <-togglec: // If paused - wait for another toggle event to unpause
	// 			<-togglec
	// 		}
	// 	}
	// }()
	// <-ui.Done()
}

// func main_bak() {
// 	var endWaiter sync.WaitGroup
// 	endWaiter.Add(1)

// 	conf, err := library.GetConf()
// 	if err != nil {
// 		log.Fatal("mainServer GetConfig Error:", err)
// 	}
// 	//flag.Parse()
// 	//logFileName := flag.String("log", conf.Setting.LogFile, "日志文件路径和名称")
// 	logFile, err := os.OpenFile(conf.Setting.LogFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("Server SetLogFile Error:", err)
// 	}
// 	log.SetOutput(logFile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// 	//获取执行参数并判断
// 	Cprocess := false
// 	Args := make([]string, len(os.Args))
// 	for i := 1; i < len(os.Args); i++ {
// 		switch os.Args[i] {
// 		//是否为子进程
// 		case "-cprocess":
// 			Cprocess = true
// 		}
// 	}
// 	//Cprocess为false则表示为父进程  判断是否需要开启后台运行
// 	if !Cprocess && conf.Setting.Daemonize == 1 {
// 		Args = append(Args, "-cprocess")
// 		// 将命令行参数中执行文件路径转换成可用路径
// 		filePath, _ := filepath.Abs(os.Args[0])
// 		cmd := exec.Command(filePath, Args...)
// 		// 将其他命令传入生成出的进程
// 		cmd.Stdin = os.Stdin // 给新进程设置文件描述符，可以重定向到文件中
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr
// 		cmd.Start() // 开始执行新进程，不等待新进程退出
// 		return
// 	}

// 	start := make(chan int)
// 	end := make(chan interface{})
// 	go func(start chan int, quit chan interface{}) {
// 		port := <-start
// 		defer recoverFromError()
// 		ui, _ := lorca.New(fmt.Sprintf("http://127.0.0.1:%d/static/index.html", port), "", 800, 600, "--disable-sync", " --disable-translate")
// 		defer ui.Close()
// 		quit <- (<-ui.Done())
// 	}(start, end)

// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	infra.ServerStart()

// 	sigs := make(chan os.Signal)
// 	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
// 	select {
// 	case <-sigs:
// 		endWaiter.Done()
// 	case <-end:
// 		endWaiter.Done()
// 	}
// 	endWaiter.Wait()
// }

// func recoverFromError() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovering from panic:", r)
// 	}
// }
