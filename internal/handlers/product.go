package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type ProductHandler struct {
	repo *repo.Repository
}

func NewProductHandler(repo *repo.Repository) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

func (h *ProductHandler) ProductHelp() {
	fmt.Println("product help")
	fmt.Println("product add            [product_name] (prod_class_id) (enum_class_id)")
	fmt.Println("product del            [product_id]")
	fmt.Println("product show           (product_id)")
	fmt.Println("product showbyprod     [prod_class_id]")
	//fmt.Println("product showbyenum     [enum_class_id]")
	//fmt.Println("product setprod        [product_id] [prod_class_id]")
	//fmt.Println("product setenum        [product_id] [enum_class_id]")
	fmt.Println("product checkcl        [product_id] [prod_class_id]")
}

func (h *ProductHandler) ProductHandle(command []string) {
	if command[0] == "help" {
		h.ProductHelp()
	} else if command[0] == "add" {
		productClassId, _ := strconv.Atoi(command[2])
		enumClassId, _ := strconv.Atoi(command[3])
		h.AddProduct(command[1], productClassId, enumClassId)
	} else if command[0] == "del" {
		productId, _ := strconv.Atoi(command[1])
		h.DeleteProduct(productId) 
	} else if command[0] == "show" {
		productId,_ := strconv.Atoi(command[1])
		h.ShowProduct(productId)
	} else if command[0] == "showbyprod" {
		classId, _ := strconv.Atoi(command[1])
		h.ShowByProdClass(classId)
	}else if command[0] == "checkcl" {//доп задание
		productId,_:=strconv.Atoi(command[1])
		classId,_:=strconv.Atoi(command[2])
		h.repo.CheckClass(context.Background(), productId, classId)
	}
}

func (h *ProductHandler) AddProduct(productName string, productClassId, enumClassId int) {
	//fmt.Println(h.repo == nil)
	h.repo.AddProduct(context.Background(), productName, productClassId, enumClassId)
}

func (h *ProductHandler) DeleteProduct(id int) {
	h.repo.DeleteProduct(context.Background(), id)
}

func (h *ProductHandler) ShowProduct (id int) {
	h.repo.ShowProduct(context.Background(), id)
}

func (h *ProductHandler) ShowByProdClass(classId int) {
	h.repo.ShowByProdClass(context.Background(), classId)
}
