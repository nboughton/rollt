// Package rollt implements a small library for creating Roll tables (commonly used in TTRPGs)
package rollt

import (
	"math/rand"
)

// Table represents a table of text options that can be rolled on. Name and Category
// are more or less optional by when dealing with large Collections can help make things
// more manageable.
type Table struct {
	Name     string
	Category string
	Opts     []Item
}

// Item represents the text and weighting (if any) of a listed item
type Item struct {
	Weight int
	Text   string
}

// Roll rolls on the table and returns the option drawn.
func (t Table) Roll() string {
	list := t.weightedList()

	return list[rand.Intn(len(list))].Text
}

func (t Table) weightedList() []Item {
	results := []Item{}
	for _, i := range t.Opts {
		for n := 0; n < i.Weight; n++ {
			results = append(results, i)
		}
	}

	return results
}

// Collection represents a group of Tables
type Collection []Table

// Names returns table Names from the Collection
func (c Collection) Names() (list []string) {
	for _, t := range c {
		list = append(list, t.Name)
	}

	return
}

// Categories returns all Category names from the Collection
func (c Collection) Categories() (list []string) {
	for _, t := range c {
		list = append(list, t.Category)
	}

	return
}
