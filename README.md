  # popup
  -使用walk-gui做的windows桌面程序
  -填入时间后到点提醒

  -编译：go build -ldflags "-H windowsgui" -o popup.exe

  编译说明文件
  -64位使用  rsrc -arch amd64 -manifest popup.manifest -o rsrc.syso
  -32位使用  rsrc -arch 386 -manifest popup.manifest -o rsrc.syso
