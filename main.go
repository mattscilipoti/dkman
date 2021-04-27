package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Generate default prompt files", "This will generate the files that create the default prompt for OPO docker projects, in docker/", 'a', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
