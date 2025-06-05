package main

import (
	"github.com/charmbracelet/bubbles/list"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

/* CUSTOM ITEM */
type Task struct {
	status      status
	title       string
	description string
}

// Implement the list.Item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

/* MAIN MODEL */
type Model struct {
	list list.Model
	err  error
}

func (m *Model) initList() {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate())
	m.list.Title = "To Do"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "Buy Eggs", description: "Free range"},
		Task{status: todo, title: "Eat Sushi", description: "Salmon skin roll"},
		Task{status: todo, title: "Pay credit card", description: "Or live in debt"},
	})
}
