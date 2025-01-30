package main

import (
	"log"

	"github.com/rivo/tview"
)

const (
	CreateProjectForm = "createProjectForm"
)

var (
	PROJECT_COMPONENTS = []string{"Message Board", "Chat", "Schedule"}
)

func main() {
	app := createApp()

	flex := tview.NewFlex()

	pages := tview.NewPages().
		AddPage("background", flex, true, true)
	pages.SendToFront("background")

	projects, err := getProjects()
	if err != nil {
		log.Fatal(err)
	}

	messages, err := getMessages()
	if err != nil {
		log.Fatal(err)
	}

	list := tview.NewList()
	flex.AddItem(list, 0, 1, true)

	list.
		SetChangedFunc(getProjectListChanged(flex, projects, messages)).
		ShowSecondaryText(false)

	for _, p := range projects {
		list.AddItem(p.Name, "", rune(0), nil)
	}

	list.SetBorder(true).
		SetTitle("Projects").
		SetInputCapture(getProjectListInputChanged(list, pages, flex, app))

	if err := app.SetRoot(pages, true).SetFocus(flex).Run(); err != nil {
		log.Fatal(err)
	}
}
