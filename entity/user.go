package entity

import (
	"strings"
	"fmt"
)

type Handler interface {
	Handle(u *User) error
}

type errorList []error

func (eL errorList) Error() string {
	var sb strings.Builder
	for i := range eL {
		sb.WriteString(eL[i].Error())
	}
	return sb.String()
}


type User struct {
	ID                   string
	Score                uint64
	MostScore            uint64
	PrevMostScore        uint64
	mostScoreChangeHandlers []Handler
}

func (u *User) OnMostScoreChange(h Handler) {
	u.mostScoreChangeHandlers = append(u.mostScoreChangeHandlers, h)
}

func (u *User) mostScoreChanged() error {
	var eL errorList
	for i, h := range u.mostScoreChangeHandlers {
		if err := h.Handle(u); err != nil {
			eL = append(eL, fmt.Errorf("Error on func %d: %w", i+1, err))
		}
	}
	return eL
}

func (u *User) IncreaseScore(s uint64) error {
	u.Score += s
	if u.Score > u.MostScore {
		u.PrevMostScore = u.MostScore
		u.MostScore = u.Score
		if err := u.mostScoreChanged(); err != nil {
			return fmt.Errorf("Error increasing user %s score: %w", u.ID, err)
		}
	}
	return nil
}

func (u *User) DecreaseScore(s uint64) {
	if u.Score <= s {
		u.Score = 0
		return
	}
	u.Score -= s
	return
}
