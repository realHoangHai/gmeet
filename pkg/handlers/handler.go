package handlers

import (
	"github.com/realHoangHai/gmeet-biz/ent"
	"github.com/realHoangHai/gmeet-biz/pkg/config"
)

type Handlers struct {
	Client *ent.Client
	Config *config.Config
}

func NewHandlers(client *ent.Client, config *config.Config) *Handlers {
	return &Handlers{
		Client: client,
		Config: config,
	}
}
