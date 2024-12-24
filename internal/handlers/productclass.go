package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type ProductClassHadnler struct {
	repo *repo.Repository
}

func NewProductClassHandler(repo *repo.Repository) *ProductClassHadnler {
	return &ProductClassHadnler{
		repo: repo,
	}
}

func (h *ProductClassHadnler) ProductClassHelp() {
	fmt.Println("prodcl help")
	fmt.Println("prodcl add         [class_name] (unit_id) (parent_id)")
	fmt.Println("prodcl del         [class_id]")
	//fmt.Println("prodcl show        (class_id)")
	fmt.Println("prodcl parent      [class_id]")
	fmt.Println("prodcl children    [class_id]")
	fmt.Println("prodcl chpar       [class_id] [parent_id]")
	fmt.Println("prodcl setunit     [class_id]")
	//fmt.Println("prodcl swap        [class_id1] [class_id2]")
}

func (h *ProductClassHadnler) ProductClassHandle(command []string) {
	if command[0] == "help" {
		h.ProductClassHelp()
	} else if command[0] == "add" {
		unitId, _ := strconv.Atoi(command[2])
		parentId, _ := strconv.Atoi(command[3])
		h.repo.AddProductClass(context.Background(), command[1], unitId, parentId)
	} else if command[0] == "delete" {

	} else if command[0] == "parent" {//показать родителей

	} else if command[0] == "children" {//показать потомков
		h.repo.PrintChildrenRecursive(context.Background())
	} else if command[0] == "chpar" {
		classId, _ := strconv.Atoi(command[1])
		newParentId, _ := strconv.Atoi(command[2])
		h.repo.ChangeParent(context.Background(), classId, newParentId)
	} else if command[0] == "setunit" {
		classId, _ := strconv.Atoi(command[1])
		unitId, _ := strconv.Atoi(command[2])
		h.repo.SetUnit(context.Background(), classId, unitId)
	} 
}
