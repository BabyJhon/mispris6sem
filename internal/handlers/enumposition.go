package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type EnumPositionHandler struct {
	repo *repo.Repository
}

func NewEnumPositionHandler(repo *repo.Repository) *EnumPositionHandler {
	return &EnumPositionHandler{
		repo: repo,
	}
}

func (h *EnumPositionHandler) EnumPositionHelp() {
	fmt.Println("enumpos help")
	fmt.Println("enumpos add       [название] [краткое_название] [ид_класса] [integer/real/string] [значение]")
	fmt.Println("enumpos delete    [ид_позиции]")
	fmt.Println("enumpos show      (ид_позиции)")
	fmt.Println("enumpos swap      [ид_позиции1] [ид_позиции2]")
}

func (h *EnumPositionHandler) EnumPositionHandle(command []string) {
	if command[0] == "help" {
		h.EnumPositionHelp()
	} else if command[0] == "add" {
		classifierId, _ := strconv.Atoi(command[3])
		h.repo.AddEnumPosition(context.Background(), command[1], command[2], classifierId, command[4], command[5])
	} else if command[0] == "delete" {
		id, _ := strconv.Atoi(command[1])
		h.repo.EnumPosition.DeleteEnumPosition(context.Background(), id)
	} else if command[0] == "show" {
		id, _ := strconv.Atoi(command[1])
		h.repo.EnumPosition.ShowEnumPosition(context.Background(), id)
	} else if command[0] == "swap" {

	}
}
