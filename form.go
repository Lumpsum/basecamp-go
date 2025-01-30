package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProjectFormOpts func(*ProjectForm)

type ProjectForm struct {
	form *tview.Form
}

func NewProjectForm(opts ...ProjectFormOpts) *ProjectForm {
	pf := &ProjectForm{}

	createProjectForm := tview.NewForm().
		AddInputField("Project Name", "", 20, nil, nil).
		AddInputField("Description", "", 20, nil, nil).
		AddButton("Add project", func() {})

	createProjectForm.
		SetFieldBackgroundColor(tcell.ColorGrey).
		SetLabelColor(tcell.ColorWhite).
		SetButtonBackgroundColor(tcell.ColorGrey).
		SetBackgroundColor(tcell.ColorBlack)

	createProjectForm.
		SetBorder(true).
		SetTitle("Create Project")

	pf.form = createProjectForm

	for _, opt := range opts {
		opt(pf)
	}

	return pf
}

func WithInputCapture(f func(event *tcell.EventKey) *tcell.EventKey) func(*ProjectForm) {
	return func(p *ProjectForm) {
		p.form.SetInputCapture(f)
	}
}
