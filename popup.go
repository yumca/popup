package main

import (
	"fmt"
	"popup/ui"
)

func main() {
	if err := ui.UiWalkWindowNew(480, 320); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ui.UiWalkWindowRun())
	ui.WalkWindowDone()
	// defer ui.LorcaClose()
	// <-ui.LorcaDone()

}
