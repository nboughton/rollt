// Package rollt implements a small library for creating Roll tables (commonly used in TTRPGs)
package rollt

import (
	"fmt"

	"github.com/nboughton/go-dice"
)

// Table represents a table of text options that can be rolled on
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
func (t *Table) Roll() string {
	list := t.weightedList()
	d, _ := dice.NewDice(fmt.Sprintf("1d%d", len(list)))
	i, _ := d.Roll()

	return list[i-1].Text
}

func (t *Table) weightedList() []Item {
	results := []Item{}
	for _, i := range t.Opts {
		for n := 0; n < i.Weight; n++ {
			results = append(results, i)
		}
	}

	return results
}

// Collection represents a group of Tables
type Collection []*Table

// Names returns the assigned table names of all tables in a collection
func (c Collection) Names() (list []string) {
	for _, t := range c {
		list = append(list, t.Name)
	}

	return
}
