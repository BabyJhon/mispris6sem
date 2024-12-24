package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type ParamProductHandler struct {
	repo *repo.Repository
}

func NewParamProductHandler(repo *repo.Repository) *ParamProductHandler {
	return &ParamProductHandler{
		repo: repo,
	}
}

func (h *ParamProductHandler) ParamProductHelp() {
	fmt.Println("parampr help")
	fmt.Println("parampr add          [product_id] [param_class_id] [value]")
	fmt.Println("parampr show         (product_id)")
	fmt.Println("parampr edit         [product_id] [param_class_id] [new_value]")
}

func (h *ParamProductHandler) ParamProductHandle(command []string) {
	if command[0] == "help" {
		h.ParamProductHelp()
	} else if command[0] == "add" {
		var productId, paramClassId, value int
		productId, _ = strconv.Atoi(command[1])
		paramClassId, _ = strconv.Atoi(command[2])
		value, _ = strconv.Atoi(command[3])
		h.repo.AddParamToProduct(context.Background(), value, paramClassId, productId)
	} else if command[0] == "show" {
		productId, _ := strconv.Atoi(command[1])
		h.repo.ShowAllByProduct(context.Background(), productId)
	} else if command[0] == "edit" {
		var productId, paramClassId, value int
		productId, _ = strconv.Atoi(command[1])
		paramClassId, _ = strconv.Atoi(command[2])
		value, _ = strconv.Atoi(command[3])
		h.repo.ParamProducts.Edit(context.Background(), productId, paramClassId, value)
	}
}
