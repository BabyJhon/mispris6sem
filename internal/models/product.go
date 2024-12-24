package models

import "fmt"

type Product struct {
	ID               int    `db:"id_product"`
	ProductName      string `db:"product_name"`
	ProductClassId   int    `db:"prod_class_id"`
	EnumClassifierId int    `db:"enum_classifier_id"`
}

func (p *Product) Repr() string {
	return fmt.Sprintf("Product(id=%d, name=%s, product_class_id=%d, enum_classifier_id=%d)", p.ID, p.ProductName, p.ProductClassId, p.EnumClassifierId)
}