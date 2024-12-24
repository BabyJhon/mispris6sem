package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type ParamHandler struct {
	repo *repo.Repository
}

func NewParamHandler(repo *repo.Repository) *ParamHandler {
	return &ParamHandler{
		repo: repo,
	}
}

func (h *ParamHandler) ParamHelp() {
	fmt.Println("param help")
	fmt.Println("param add          [name] [short_name] (unit_id) (enum_cl_id)")
}

func (h *ParamHandler) ParamHandle(command []string) {
	if command[0] == "help" {
		h.ParamHelp()
	} else if command[0] == "add" {
		unitId, _ := strconv.Atoi(command[3])
		enumClassifierId, _:= strconv.Atoi(command[4])
		h.repo.AddParam(context.Background(), command[1], command[2], unitId, enumClassifierId)
	}
}
