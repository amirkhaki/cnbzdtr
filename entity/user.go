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
	ID                   string `redis:"-"`
	Score                uint64 `redis:"score"`
	MostScore            uint64 `redis:"most_score"`
	PrevMostScore        uint64 `redis:"prev_most_score"`
	mostScoreChangeHandlers map[string]Handler
	ReferralCount uint64 `redis:"referral_count"`
}

func (u *User) OnMostScoreChange(title string, h Handler) {
	u.mostScoreChangeHandlers[title] = h
}

func (u *User) mostScoreChanged() error {
	var eL errorList
	for i, h := range u.mostScoreChangeHandlers {
		if err := h.Handle(u); err != nil {
			eL = append(eL, fmt.Errorf("Error on func %s: %w", i, err))
		}
	}
	if len(eL) == 0 {
		return nil
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

func (u *User) AddReferral(referredID string) {
	u.ReferralCount += 1
}

func NewUser(id string) (*User, error) {
	u := User{ID:id}
	u.mostScoreChangeHandlers = make(map[string]Handler)
	return &u, nil
}
