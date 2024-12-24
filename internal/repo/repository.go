package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EnumClassifier interface {
	AddEnumClassifier(ctx context.Context, name string, parentId int)
	DeleteEnumClassifier(ctx context.Context, id int)
}

type EnumPosition interface {
	AddEnumPosition(ctx context.Context, name, shortName string, classifierId int, valueType string, value string)
	DeleteEnumPosition(ctx context.Context, id int)
	ShowEnumPosition(ctx context.Context, id int)
}

type Param interface {
	AddParam(ctx context.Context, name, shortName string, unitId, enumClassifierId int)
}

type ParamClass interface {
	AddParamToClass(ctx context.Context, paramId, prodclassId, minValue, maxValue int)
	ShowByClass(ctx context.Context, prodClassId int)
}

type ParamProducts interface {
	AddParamToProduct(ctx context.Context, value, paramClassId, productId int)
	Edit(ctx context.Context, productId, paramClassId, value int)
	ShowAllByProduct(ctx context.Context, productId int)
}

type ProdClass interface {
	AddProductClass(ctx context.Context, className string, unitId, parentId int)
	ChangeParent(ctx context.Context, classId, newParentId int)
	SetUnit(ctx context.Context, classId, unitId int)
	PrintChildrenRecursive(ctx context.Context)
}

type Product interface {
	AddProduct(ctx context.Context, productName string, productClassId, enumClassId int)
	DeleteProduct(ctx context.Context, id int)
	ShowProduct(ctx context.Context, id int)
	ShowByProdClass(ctx context.Context, classId int)
	CheckClass(ctx context.Context, productId, classId int)
}

type Unit interface {
	AddUnit(ctx context.Context, unitName, shortName string) (int, error)
	DeleteUnit(ctx context.Context, id int)
	ShowUnit(ctx context.Context, id int)
	UpdateUnit(ctx context.Context, id int, name, shortName string)
}

type Repository struct {
	EnumClassifier
	EnumPosition
	Param
	ParamClass
	ParamProducts
	ProdClass
	Product
	Unit
}

func NewRepositiry(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Unit:           NewUnitRepo(pool),
		Product:        NewProductRepo(pool),
		ProdClass:      NewProductClassRepo(pool),
		EnumClassifier: NewEnumClassifierRepo(pool),
		EnumPosition:   NewEnumPositionRepo(pool),
		Param:          NewParamRepo(pool),
		ParamClass:     NewParamClassRepo(pool),
		ParamProducts:  NewParamProductRepo(pool),
	}
}
