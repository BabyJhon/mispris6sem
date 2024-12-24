package handlers

import (
	"github.com/BabyJhon/mispris1-2/internal/repo"
)

type Unit interface {
	UnitHelp()
	UnitHandle(command []string)
	AddUnit(name, shortName string)
	DeleteUnit(id int)
	ShowUnit(id int)
}

type Product interface {
	AddProduct(productName string, productClassId, enumClassId int)
	ProductHandle(command []string)
	DeleteProduct(id int)
	ShowProduct(id int)
	ShowByProdClass(classId int)
}

type ProductClass interface {
	ProductClassHelp()
	ProductClassHandle(command []string)
}

type EnumClassifier interface {
	Handle(command []string)
	EnumClassifierHelp()
}

type EnumPosition interface {
	EnumPositionHelp()
	EnumPositionHandle(command []string)
}

type Param interface {
	ParamHandle(command []string)
	ParamHelp()
}

type ParamProduct interface {
	ParamProductHandle(command []string)
	ParamProductHelp() 
}

type ParamClass interface {
	ParamClassHandle(command []string)
	ParamClassHelp() 
}

type Handlers struct {
	Param
	Unit
	Product
	ProductClass
	EnumClassifier
	EnumPosition
	ParamProduct
	ParamClass
}

func NewHandlers(repo *repo.Repository) *Handlers {
	return &Handlers{
		Unit:           NewUnitHandler(repo),
		Product:        NewProductHandler(repo),
		ProductClass:   NewProductClassHandler(repo),
		EnumClassifier: NewEnumClassifierHandler(repo),
		EnumPosition:   NewEnumPositionHandler(repo),
		Param:          NewParamHandler(repo),
		ParamProduct:   NewParamProductHandler(repo),
		ParamClass:     NewParamClassHandler(repo),
	}
}
