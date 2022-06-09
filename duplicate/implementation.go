package duplicate

import (
	"fmt"
)

// Item is the type describing the duplicate
// X and Y fields are row and column, Value is the value of the duplicate
type Item struct {
	X, Y  int
	Value string
}

// Finder - interface for searching duplicates
// duplicates should be searched in columns
type Finder interface {
	Find(chan []string) []Item
}

// NewFinder is a constructor for Finder
// We should implement the Finder interface and add the body of the constructor
func NewFinder() Finder {
	return mapFinder{}
}

var _ Finder = (*mapFinder)(nil)

type mapFinder struct{}

// Find finds duplicate row values by columns
func (f mapFinder) Find(c chan []string) []Item {
	var dup []Item
	var curr int

	// uniqItem - an auxiliary structure for tracking the original take
	type uniqItem struct {
		item Item
		recv bool
	}

	cols := make(map[string]uniqItem)

	for line := range c {
		for idx := range line {
			k := fmt.Sprintf("%s:%d", line[idx], idx)

			s, ok := cols[k]
			if !ok {
				cols[k] = uniqItem{
					item: Item{
						X:     idx,
						Y:     curr,
						Value: line[idx],
					},
				}
				continue
			}

			if !s.recv {
				dup = append(dup, s.item)
				s.recv = true
				cols[k] = s
			}

			dup = append(
				dup, Item{
					X:     idx,
					Y:     curr,
					Value: line[idx],
				},
			)
		}

		curr++
	}

	return dup
}
