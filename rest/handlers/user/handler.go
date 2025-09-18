package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	userRepo repo.UserRepo
	cnf *config.Config
}

func NewHandler(cnf *config.Config, userRepo repo.UserRepo) *Handler {
	return &Handler{
		cnf: cnf,
		userRepo: userRepo,
	}
}