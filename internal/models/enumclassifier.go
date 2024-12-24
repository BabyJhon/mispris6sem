package models

type EnumClassifier struct {
	ID int `db:"id"`
	Name string `db:"name"`
	ParentID int  `db:"parent_id"`
	UnitID int  `db:"unit_id"`
	}