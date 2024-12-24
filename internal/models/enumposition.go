package models

import "fmt"

type EnumPosition struct {
	Id           int     `db:"id"`
	Name         string  `db:"name"`
	ShortName    string  `db:"short_name"`
	IntegerValue int     `db:"integer_value"`
	RealValue    float64 `db:"real_value"`
	StringValue  string  `db:"string_value"`
	ClassifierId int     `db:"classifier_id"`
}

func (ep *EnumPosition) Repr() string{
	return fmt.Sprintf("Enum Position(id=%d, name=%s, short_name=%s, integer_value=%d, real_value=%f, string_value=%s, classifier_id=%d)",
		ep.Id, ep.Name, ep.ShortName, ep.IntegerValue, ep.RealValue, ep.StringValue, ep.ClassifierId)
}
