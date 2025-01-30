package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createApp() *tview.Application {
	app := tview.NewApplication()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// switch event.Key() {
		// case tcell.KeyRune:
		// 	switch event.Rune() {
		// 	case 'q':
		//
		// 	}
		// case tcell.KeyCtrlC:
		// 	return nil
		// }
		return event
	})
	return app
}
