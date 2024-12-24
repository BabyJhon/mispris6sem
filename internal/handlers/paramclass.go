package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type ParamClassHandler struct {
	repo *repo.Repository
}

func NewParamClassHandler(repo *repo.Repository) *ParamClassHandler {
	return &ParamClassHandler{
		repo: repo,
	}
}

func (p *ParamClassHandler) ParamClassHelp() {
	fmt.Println("paramcl help")
	fmt.Println("paramcl add          [param_id] [prodclass_id] (min_v) (max_v)")
	fmt.Println("paramcl showbyclass  [prodclass_id]")
}

func (p *ParamClassHandler) ParamClassHandle(command []string) {
	if command[0] == "help" {
		p.ParamClassHelp()
	} else if command[0] == "add" {
		var paramId, prodClassId, minValue, maxValue int
		paramId, _ = strconv.Atoi(command[1])
		prodClassId, _ = strconv.Atoi(command[2])
		minValue, _ = strconv.Atoi(command[3])
		maxValue, _ = strconv.Atoi(command[4])
		p.repo.AddParamToClass(context.Background(), paramId, prodClassId, minValue, maxValue)
	} else if command[0] == "showbyclass" {
		prodclassId, _ := strconv.Atoi(command[1])
		p.repo.ShowByClass(context.Background(), prodclassId)
	}
}
