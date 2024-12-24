package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type EnumClassifierHandler struct {
	repo *repo.Repository
}

func NewEnumClassifierHandler(repo *repo.Repository) *EnumClassifierHandler {
	return &EnumClassifierHandler{
		repo: repo,
	}
}

func (h *EnumClassifierHandler) EnumClassifierHelp() {
	fmt.Println("enumcl help")
	fmt.Println("enumcl add          [название] (идентификатор_родителя)")
	fmt.Println("enumcl delete       [идентификатор_класса]")
	fmt.Println("enumcl show         (идентификатор_класса)")
	fmt.Println("enumcl parent       [идентификатор_класса]")
	fmt.Println("enumcl children     [идентификатор_класса]")
	fmt.Println("enumcl positions    [идентификатор_класса]")
	fmt.Println("enumcl chpar        [идентификатор_класса] [новый_идентификатор_родителя]")
	fmt.Println("enumcl swap         [идентификатор_класса1] [идентификатор_класса2]")
}

func (h *EnumClassifierHandler) Handle(command []string) {
	if command[0] == "help" {
		h.EnumClassifierHelp()
	} else if command[0] == "add" {
		parentId, _ := strconv.Atoi(command[2])
		h.repo.AddEnumClassifier(context.Background(), command[1], parentId)
	} else if command[0] == "delete" {

	} else if command[0] == "show" {

	} else if command[0] == "parent" {

	} else if command[0] == "children" {

	} else if command[0] == "position" {

	} else if command[0] == "chpar" {
		classId, _ := strconv.Atoi(command[1])
		newParentId, _ := strconv.Atoi(command[2])
		h.repo.ChangeParent(context.Background(), classId, newParentId)
	} else if command[0] == "swap" {

	}
}
