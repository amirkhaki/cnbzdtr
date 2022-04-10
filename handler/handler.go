package handler

import (
	"github.com/amirkhaki/cnbzdtr/config"
	"github.com/amirkhaki/cnbzdtr/entity"
	"github.com/amirkhaki/cnbzdtr/protocol"
)

type Handler struct {
	lvls  *entity.Levels
	store protocol.Store
	cfg   config.Config
}

func New(lvls *entity.Levels, s protocol.Store, cfg config.Config) *Handler {
	h := &Handler{lvls: lvls, store: s, cfg: cfg}
	return h
}
