package models

type ProductClass struct {
	ID int `db:"id"`
	ClassName string `db:"class_name"`
	UnitId int `db:"unit_id"`
	ParentId int `db:"parent_id"`
}