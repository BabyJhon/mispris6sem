package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type UnitHandler struct {
	repo *repo.Repository
}

func NewUnitHandler(repo *repo.Repository) *UnitHandler {
	return &UnitHandler{
		repo: repo,
	}
}

func (h *UnitHandler) UnitHelp() {
	fmt.Println("unit help")
	fmt.Println("unit add     [unit_name] [short_name]")
	fmt.Println("unit del     [unit_id]")
	fmt.Println("unit show    (unit_id)")
	fmt.Println("unit update  [init_id] [unit_name] [short_name]")
}

func (h *UnitHandler) UnitHandle(command []string) {

	if command[0] == "help" {
		h.UnitHelp()
	} else if command[0] == "add" {
		h.AddUnit(command[1], command[2]) //name, short_name
	} else if command[0] == "delete" {
		s, _ := strconv.Atoi(command[1])
		h.DeleteUnit(s)
	} else if command[0] == "show" {
		s, _ := strconv.Atoi(command[1])
		h.ShowUnit(s)
	} else if command[0] == "update" {
		id, _ := strconv.Atoi(command[1])
		h.UpdateUnit(id, command[2], command[3])
	} else {
		fmt.Println("unknown command")
	}
}

func (h *UnitHandler) AddUnit(name, shortName string) {
	h.repo.AddUnit(context.Background(), name, shortName)
}

func (h *UnitHandler) DeleteUnit(id int) {
	h.repo.DeleteUnit(context.Background(), id)
}

func (h *UnitHandler) ShowUnit(id int) {
	h.repo.ShowUnit(context.Background(), id)
}

func (h *UnitHandler) UpdateUnit(id int, name, shortName string) {
	h.repo.UpdateUnit(context.Background(), id, name, shortName)
}
