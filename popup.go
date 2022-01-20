package main

import "popup/ui"

func main() {
	ui.UiWalkWindowNew(480, 320)
	ui.UiWalkWindowRun()
	// defer ui.LorcaClose()
	// <-ui.LorcaDone()

}
