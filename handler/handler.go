package handler

import (
	"github.com/amirkhaki/cnbzdtr/entity"
	"github.com/amirkhaki/cnbzdtr/protocol"
)

type Handler struct {
	lvls *entity.Levels
	store protocol.Store
}

func New(lvls *entity.Levels, s protocol.Store) *Handler {
	h := &Handler{lvls:lvls, store: s}
	return h
}
