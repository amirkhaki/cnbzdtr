package handler

import (
	"context"
	dg "github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) InviteCreate(ctx context.Context, s *dg.Session, i *dg.InviteCreate) {
	user, err := h.store.GetUserOrCreate(ctx, i.Inviter.ID)
	if err != nil {
		log.Println(err)
		return
	}
	seh := &SEH{s: s, lvls: h.lvls, cfg: h.cfg}
	user.OnMostScoreChange("SEH", seh)
	user.AddReferral(i.TargetUser.ID)
	err = user.IncreaseScore(h.cfg.InviteScore)
	if err != nil {
		log.Println(err)
		return
	}
	err = h.store.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err)
	}
}
