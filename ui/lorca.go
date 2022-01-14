package ui

import (
	"log"
	"popup/library"
	"popup/view"

	"github.com/yumca/lorca"
)

var (
	Lorcaui lorca.UI
	path    string
	err     error
)

func LorcaNew(width, height int, customArgs ...string) error {
	path = library.GetExecPath()
	Lorcaui, err = lorca.New(view.GetView(""), "", width, height, customArgs...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	jsbind()
	frameTree, _ := Lorcaui.GetFrameTree()
	Lorcaui.SetContent(frameTree.FrameTree.Frame.Id, view.GetView("main"))
	Lorcaui.SetDebugger(false)
	return nil
}

func LorcaClose() {
	Lorcaui.Close()
}

func LorcaDone() <-chan struct{} {
	return Lorcaui.Done()
}
