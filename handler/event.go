package handler

import (
	dg "github.com/bwmarrin/discordgo"
	"github.com/amirkhaki/cnbzdtr/entity"
	"fmt"
)

// score change event handler
type SEH struct {
	s *dg.Session
	lvls *entity.Levels
}
func (sh *SEH) Handle(u *entity.User) error {
	crrntLevel := sh.lvls.Level(u.MostScore)
	prevLevel := sh.lvls.Level(u.PrevMostScore)
	if crrntLevel == prevLevel {
		return nil
	}
	ch, err := sh.s.UserChannelCreate(u.ID)
	if err != nil {
		return fmt.Errorf("Could not create DM channel with user: %w", err)
	}
	message := fmt.Sprintf("You were at level %s, and now are at %s",prevLevel.Title,crrntLevel.Title)
	_, err = sh.s.ChannelMessageSend(ch.ID, message)
	if err != nil {
		return fmt.Errorf("Could not send DM to user: %w", err)
	}
	return nil
}
