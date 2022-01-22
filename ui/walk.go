package ui

import (
	"popup/model/tables"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

var UiMainWindow = new(UiWalkWindow)

var lb *walk.ListBox
var model *contentModel
var input *walk.TextEdit

type UiWalkWindow struct {
	*walk.MainWindow
}

func UiWalkWindowNew(width, height int) error {
	//列表单条内容结构
	var items []contentEntry
	mettings := get_meetings()
	// var cstSh = time.FixedZone("CST", 8*3600) //东八区
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
	//列表内容模型
	model = &contentModel{items: items}
	//设置列表风格
	styler := &Styler{
		lb:                  &lb,
		model:               model,
		dpi2StampSize:       make(map[int]walk.Size),
		widthDPI2WsPerLine:  make(map[widthDPI]int),
		textWidthDPI2Height: make(map[textWidthDPI]int),
	}

	if err := (MainWindow{
		//引入窗口
		AssignTo: &UiMainWindow.MainWindow,
		Title:    "会议提醒程序",
		MenuItems: []MenuItem{
			Menu{
				Text: "操作",
				Items: []MenuItem{
					Action{
						Text:        "刷新",
						OnTriggered: reflash,
					},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { UiMainWindow.Close() },
					},
				},
			},
		},
		//窗口最小大小
		MinSize: Size{Width: 200, Height: 200},
		//设置窗口大小
		Size:   Size{Width: width, Height: height},
		Font:   Font{Family: "Segoe UI", PointSize: 9},
		Layout: VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &input,
				MaxSize:  Size{Width: width - 24, Height: height / 2},
			},
			PushButton{
				Text:      "add",
				OnClicked: save_meeting,
			},
			Composite{
				DoubleBuffering: true,
				Layout:          VBox{},
				Children: []Widget{
					ListBox{
						AssignTo:       &lb,
						MultiSelection: true,
						Model:          model,
						ItemStyler:     styler,
					},
				},
			},
			PushButton{
				Text:      "delete",
				OnClicked: delete_meeting,
			},
		},
	}.Create()); err != nil {
		return err
	}
	notifyTicker()

	return nil
}

func UiWalkWindowRun() int {
	UiMainWindow.Show()
	return UiMainWindow.Run()
}

func WalkWindowDone() {
	cancel <- true
}

func (uw *UiWalkWindow) msgBoxTriggered() {
	walk.MsgBox(uw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

type contentModel struct {
	walk.ReflectListModelBase
	items []contentEntry
}

func (m *contentModel) Items() interface{} {
	return m.items
}

type contentEntry struct {
	id        int
	timestamp string
	content   string
	notify    string
}

type widthDPI struct {
	width int // in native pixels
	dpi   int
}

type textWidthDPI struct {
	text  string
	width int // in native pixels
	dpi   int
}
type Styler struct {
	lb                  **walk.ListBox
	canvas              *walk.Canvas
	model               *contentModel
	font                *walk.Font
	dpi2StampSize       map[int]walk.Size
	widthDPI2WsPerLine  map[widthDPI]int
	textWidthDPI2Height map[textWidthDPI]int // in native pixels
}

func (s *Styler) ItemHeightDependsOnWidth() bool {
	return true
}

func (s *Styler) DefaultItemHeight() int {
	dpi := (*s.lb).DPI()
	marginV := walk.IntFrom96DPI(marginV96dpi, dpi)

	return s.StampSize().Height + marginV*2
}

const (
	marginH96dpi int = 6
	marginV96dpi int = 2
	lineW96dpi   int = 1
)

func (s *Styler) ItemHeight(index, width int) int {
	dpi := (*s.lb).DPI()
	marginH := walk.IntFrom96DPI(marginH96dpi, dpi)
	marginV := walk.IntFrom96DPI(marginV96dpi, dpi)
	lineW := walk.IntFrom96DPI(lineW96dpi, dpi)

	msg := s.model.items[index].content

	twd := textWidthDPI{msg, width, dpi}

	if height, ok := s.textWidthDPI2Height[twd]; ok {
		return height + marginV*2
	}

	canvas, err := s.Canvas()
	if err != nil {
		return 0
	}

	notifySize := s.NotifySize()

	wd := widthDPI{width, dpi}
	wsPerLine, ok := s.widthDPI2WsPerLine[wd]
	if !ok {
		bounds, _, err := canvas.MeasureTextPixels("W", (*s.lb).Font(), walk.Rectangle{Width: 9999999}, walk.TextCalcRect)
		if err != nil {
			return 0
		}
		wsPerLine = (width - marginH*4 - lineW - notifySize.Width) / bounds.Width
		s.widthDPI2WsPerLine[wd] = wsPerLine
	}

	if len(msg) <= wsPerLine {
		s.textWidthDPI2Height[twd] = notifySize.Height
		return notifySize.Height + marginV*2
	}

	bounds, _, err := canvas.MeasureTextPixels(msg, (*s.lb).Font(), walk.Rectangle{Width: width - marginH*4 - lineW - notifySize.Width, Height: 255}, walk.TextEditControl|walk.TextWordbreak|walk.TextEndEllipsis)
	if err != nil {
		return 0
	}

	s.textWidthDPI2Height[twd] = bounds.Height

	return bounds.Height + marginV*2
}

func (s *Styler) StyleItem(style *walk.ListItemStyle) {
	if canvas := style.Canvas(); canvas != nil {
		if style.Index()%2 == 1 && style.BackgroundColor == walk.Color(win.GetSysColor(win.COLOR_WINDOW)) {
			style.BackgroundColor = walk.Color(win.GetSysColor(win.COLOR_BTNFACE))
			if err := style.DrawBackground(); err != nil {
				return
			}
		}

		pen, err := walk.NewCosmeticPen(walk.PenSolid, style.LineColor)
		if err != nil {
			return
		}
		defer pen.Dispose()

		dpi := (*s.lb).DPI()
		marginH := walk.IntFrom96DPI(marginH96dpi, dpi)
		marginV := walk.IntFrom96DPI(marginV96dpi, dpi)
		lineW := walk.IntFrom96DPI(lineW96dpi, dpi)

		b := style.BoundsPixels()
		b.X += marginH
		b.Y += marginV

		item := s.model.items[style.Index()]

		style.DrawText(item.notify, b, walk.TextEditControl|walk.TextWordbreak)

		stampSize := s.StampSize()

		x := b.X + stampSize.Width + marginH + lineW
		canvas.DrawLinePixels(pen, walk.Point{x, b.Y - marginV}, walk.Point{x, b.Y - marginV + b.Height})

		b.X += stampSize.Width + marginH*2 + lineW
		b.Width -= stampSize.Width + marginH*4 + lineW

		style.DrawText(item.timestamp, b, walk.TextEditControl|walk.TextWordbreak|walk.TextEndEllipsis)

		notifySize := s.NotifySize()

		x = b.X + notifySize.Width + marginH + lineW
		canvas.DrawLinePixels(pen, walk.Point{x, b.Y - marginV}, walk.Point{x, b.Y - marginV + b.Height})

		b.X += notifySize.Width + marginH*2 + lineW
		b.Width -= notifySize.Width + marginH*4 + lineW

		style.DrawText(item.content, b, walk.TextEditControl|walk.TextWordbreak|walk.TextEndEllipsis)
	}
}

func (s *Styler) StampSize() walk.Size {
	dpi := (*s.lb).DPI()

	stampSize, ok := s.dpi2StampSize[dpi]
	if !ok {
		canvas, err := s.Canvas()
		if err != nil {
			return walk.Size{}
		}

		bounds, _, err := canvas.MeasureTextPixels("Jan _2 20:04:05.000", (*s.lb).Font(), walk.Rectangle{Width: 9999999}, walk.TextCalcRect)
		if err != nil {
			return walk.Size{}
		}

		stampSize = bounds.Size()
		s.dpi2StampSize[dpi] = stampSize
	}

	return stampSize
}

func (s *Styler) NotifySize() walk.Size {
	dpi := (*s.lb).DPI()

	stampSize, ok := s.dpi2StampSize[dpi]
	if !ok {
		canvas, err := s.Canvas()
		if err != nil {
			return walk.Size{}
		}

		bounds, _, err := canvas.MeasureTextPixels("已通知", (*s.lb).Font(), walk.Rectangle{Width: 9999999}, walk.TextCalcRect)
		if err != nil {
			return walk.Size{}
		}

		stampSize = bounds.Size()
		s.dpi2StampSize[dpi] = stampSize
	}

	return stampSize
}

func (s *Styler) Canvas() (*walk.Canvas, error) {
	if s.canvas != nil {
		return s.canvas, nil
	}

	canvas, err := (*s.lb).CreateCanvas()
	if err != nil {
		return nil, err
	}
	s.canvas = canvas
	(*s.lb).AddDisposable(canvas)

	return canvas, nil
}
