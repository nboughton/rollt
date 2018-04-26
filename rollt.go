// Package rollt implements a small library for creating Roll tables (commonly used in TTRPGs)
package rollt

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/nboughton/go-dice"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Table represents a table of text options that can be rolled on. Name is
// optional. Tables are preferable to Lists when using multiple dice to achieve
// a result (i.e 2d6) because their results fall on a bell curve whereas single-die
// rolls have an even probability.
type Table struct {
	Name  string
	Dice  string
	Items []Item
}

// Item represents the text and matching numbers from the table
type Item struct {
	Match []int
	Text  string
}

// Roll on the table and return the option drawn.
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

func (t Table) String() string {
	var (
		buf = new(bytes.Buffer)
		tw  = tabwriter.NewWriter(buf, 2, 2, 1, ' ', 0)
	)

	fmt.Fprintln(tw, "Dice\t|\tText")
	for _, i := range t.Items {
		fmt.Fprintf(tw, "%s\t|\t%s\n", matchToStr(i.Match), i.Text)
	}
	tw.Flush()

	return buf.String()
}

func matchToStr(m []int) string {
	var s []string
	for _, n := range m {
		s = append(s, strconv.Itoa(n))
	}

	return strings.Join(s, ", ")
}

// List represents a List of strings from which something can be selected at random
type List struct {
	Name  string
	Items []string
}

// Roll returns a random string from List
func (l List) Roll() string {
	if len(l.Items) > 0 {
		return l.Items[rand.Intn(len(l.Items))]
	}

	return ""
}

func (l List) String() string {
	return strings.Join(l.Items, ", ")
}
