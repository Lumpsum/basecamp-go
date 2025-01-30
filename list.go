package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setBasicListNavigation(list *tview.List, event *tcell.EventKey) {
	itemAmount := list.GetItemCount()
	newItem := list.GetCurrentItem()
	if event.Rune() == 'j' {
		newItem = newItem + 1
		if newItem >= itemAmount {
			newItem = 0
		}
	} else {
		newItem = newItem - 1
		if newItem < 0 {
			newItem = itemAmount
		}
	}
	list.SetCurrentItem(newItem)
}

func getProjectListChanged(flex *tview.Flex, projects []Project, messages []Message) func(i int, _, _ string, _ rune) {
	return func(i int, _, _ string, _ rune) {
		if flex.GetItemCount() > 1 {
			flex.RemoveItem(flex.GetItem(flex.GetItemCount() - 1))
		}

		innerFlex := tview.NewFlex().SetDirection(tview.FlexRow)
		for _, d := range projects[i].Dock {
			if d.Enabled {
				switch d.Title {
				case "Message Board":
					messageList := tview.NewList()
					for _, m := range messages {
						messageList.AddItem(m.Title, "", rune(0), nil)
					}
					messageList.ShowSecondaryText(false)

					messageBoard := tview.NewFlex()
					messageBoard.AddItem(messageList, 0, 1, false)
					messageBoard.SetTitle(d.Title).
						SetBorder(true)

					innerFlex.AddItem(
						messageBoard,
						0,
						1,
						false,
					)
				}
			}
		}
		innerFlex.SetTitle(projects[i].Name).SetBorder(true)
		flex.AddItem(innerFlex, 0, 5, false)
	}
}

func getProjectListInputChanged(list *tview.List, pages *tview.Pages, flex *tview.Flex, app *tview.Application) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j', 'k':
			setBasicListNavigation(list, event)
		case 'a':
			projectForm := NewProjectForm(
				WithInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
					switch event.Rune() {
					case 'q':
						pages.RemovePage(CreateProjectForm)
						app.SetFocus(flex)
						return nil
					}
					return event
				}),
			)
			pages.AddPage(CreateProjectForm, getCenteredModal(projectForm.form, 40, 10), true, true)
			app.SetFocus(projectForm.form)
		}

		return event
	}
}
