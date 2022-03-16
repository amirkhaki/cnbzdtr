package entity

import (
	"sort"
)

type Level struct {
	Title string
	From  uint64
}

type Levels struct {
	list []Level
}

// add level to levels list
func (l *Levels) AddLevel(lvl Level) {
	l.list = append(l.list, lvl)
}

// sort levels
func (l *Levels) Sort() {
	sort.Slice(l.list, func(i, j int) bool {
		return l.list[i].From < l.list[j].From
	})
}

// find what level score is in
func (l *Levels) Level(s uint64) Level {
	if len(l.list) == 0 {
		return Level{Title: "No level exists!"}
	}
	l.Sort()
	for i := range l.list {
		if i+1 == len(l.list) {
			if s >= l.list[i].From {
				return l.list[i]
			}
		}
		if s >= l.list[i].From && s < l.list[i+1].From {
			return l.list[i]
		}
	}
	return l.list[0]
}
