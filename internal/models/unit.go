package models

import "fmt"

type Unit struct {
	ID         int `db:"id"`
	Unit_name  string `db:"unit_name"`
	Short_name string `db:"short_name"`
}

func (u *Unit) Repr() string {
	return fmt.Sprintf("Unit(id=%d, unit_name=%s, short_name=%s)", u.ID, u.Unit_name, u.Short_name)
}