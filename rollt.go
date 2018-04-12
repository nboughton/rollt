// Package rollt implements a small library for creating Roll tables (commonly used in TTRPGs)
package rollt

import (
	"math/rand"
	"time"

	"github.com/nboughton/go-dice"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Table represents a table of text options that can be rolled on. Name and Category
// are more or less optional by when dealing with large Collections can help make things
// more manageable.
type Table struct {
	Name     string
	Category string
	Dice     string
	Items    []Item
}

// List represents a List of strings from which something can be selected at random
type List []string

// Roll returns a random string from List
func (l List) Roll() string {
	return l[rand.Intn(len(l))]
}

// Item represents the text and matching numbers from the table
type Item struct {
	Match []int
	Text  string
}

// Roll rolls on the table and returns the option drawn.
func (t Table) Roll() string {
	d, err := dice.NewDice(t.Dice)
	if err != nil {
		panic(err)
	}

	n, _ := d.Roll()

	for _, i := range t.Items {
		for _, m := range i.Match {
			if m == n {
				return i.Text
			}
		}
	}

	return ""
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
