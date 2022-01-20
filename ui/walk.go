package ui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()
var UiMainWindow = new(UiWalkWindow)

type UiWalkWindow struct {
	*walk.MainWindow
}

func UiWalkWindowNew(width, height int, customArgs ...string) error {
	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu
	var toggleSpecialModePB *walk.PushButton
	if err := (MainWindow{
		AssignTo: &UiMainWindow.MainWindow,
		Title:    "Walk Actions Example",
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "../img/open.png",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: UiMainWindow.openAction_Triggered,
					},
					Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { UiMainWindow.Close() },
					},
				},
			},
			Menu{
				Text: "&View",
				Items: []MenuItem{
					Action{
						Text:    "Open / Special Enabled",
						Checked: Bind("enabledCB.Visible"),
					},
					Action{
						Text:    "Open Hidden",
						Checked: Bind("openHiddenCB.Visible"),
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: UiMainWindow.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				ActionRef{&openAction},
				Menu{
					Text:  "New A",
					Image: "../img/document-new.png",
					Items: []MenuItem{
						Action{
							Text:        "A",
							OnTriggered: UiMainWindow.newAction_Triggered,
						},
						Action{
							Text:        "B",
							OnTriggered: UiMainWindow.newAction_Triggered,
						},
						Action{
							Text:        "C",
							OnTriggered: UiMainWindow.newAction_Triggered,
						},
					},
					OnTriggered: UiMainWindow.newAction_Triggered,
				},
				Separator{},
				Menu{
					Text:  "View",
					Image: "../img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text:        "X",
							OnTriggered: UiMainWindow.changeViewAction_Triggered,
						},
						Action{
							Text:        "Y",
							OnTriggered: UiMainWindow.changeViewAction_Triggered,
						},
						Action{
							Text:        "Z",
							OnTriggered: UiMainWindow.changeViewAction_Triggered,
						},
					},
				},
				Separator{},
				Action{
					Text:        "Special",
					Image:       "../img/system-shutdown.png",
					Enabled:     Bind("isSpecialMode && enabledCB.Checked"),
					OnTriggered: UiMainWindow.specialAction_Triggered,
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			CheckBox{
				Name:    "enabledCB",
				Text:    "Open / Special Enabled",
				Checked: true,
				Accessibility: Accessibility{
					Help: "Enables Open and Special",
				},
			},
			CheckBox{
				Name:    "openHiddenCB",
				Text:    "Open Hidden",
				Checked: true,
			},
			PushButton{
				AssignTo: &toggleSpecialModePB,
				Text:     "Enable Special Mode",
				OnClicked: func() {
					isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())

					if isSpecialMode.Satisfied() {
						toggleSpecialModePB.SetText("Disable Special Mode")
					} else {
						toggleSpecialModePB.SetText("Enable Special Mode")
					}
				},
				Accessibility: Accessibility{
					Help: "Toggles special mode",
				},
			},
		},
	}.Create()); err != nil {
		return err
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(UiMainWindow.openAction_Triggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	return nil
}

func UiWalkWindowRun() int {
	return UiMainWindow.Run()
}

func UiWalkWindowClose() {
}

func WalkWindowDone() <-chan struct{} {
	return nil
}

func (mw *UiWalkWindow) openAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *UiWalkWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *UiWalkWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *UiWalkWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *UiWalkWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
