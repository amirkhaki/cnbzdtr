package entity

import (
	"strings"
)

type User struct {
	ID                   string
	Score                uint64
	MostScore            uint64
	PrevMostScore        uint64
	mostScoreChangeFuncs []func(u *User) error
}

func (u *User) OnMostScoreChange(f func(u *User) error) {
	u.mostScoreChangeFuncs = append(u.mostScoreChangeFuncs, f)
}

type errorList []error

func (eL errorList) Error() string {
	var sb strings.Builder
	for i := range eL {
		sb.WriteString(eL[i].Error())
	}
	return sb.String()
}
func (u *User) mostScoreChanged() error {
	var eL errorList
	for i, f := range u.mostScoreChangeFuncs {
		if err := f(u); err != nil {
			eL = append(eL, fmt.Errorf("Error on func %d: %w", i+1, err))
		}
	}
	return eL
}
