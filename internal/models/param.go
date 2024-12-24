package models

type Param struct {
	// id = Column(Integer, primary_key=True)
	// name = Column(String(100), unique=True)
	// short_name = Column(String(20))
	// unit_id = Column(Integer, ForeignKey('unit.id'))
	// enum_classifier_id = Column(Integer, ForeignKey('enum_classifier.id'))
	Id               int    `db:"id"`
	Name             string `db:"name"`
	ShortName        string `db:"short_name"`
	UnitId           int    `db:"unit_id"`
	EnumClassifierId int    `db:"enum_classifier_id"`
}
