package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/BabyJhon/mispris1-2/internal/models"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EnumPositionRepo struct {
	db *pgxpool.Pool
}

func NewEnumPositionRepo(db *pgxpool.Pool) *EnumPositionRepo {
	return &EnumPositionRepo{
		db: db,
	}
}

func (r *EnumPositionRepo) AddEnumPosition(ctx context.Context, name, shortName string, classifierId int, valueType string, value string) {
	var id int

	var query string
	if valueType == "integer" {
		query = fmt.Sprintf("INSERT INTO enum_position (name, short_name, classifier_id, integer_value) values($1,$2,$3,$4) RETURNING id")
	} else if valueType == "string" {
		query = fmt.Sprintf("INSERT INTO enum_position (name, short_name, classifier_id, string_value) values($1,$2,$3,$4) RETURNING id")
	} else if valueType == "real" {
		query = fmt.Sprintf("INSERT INTO enum_position (name, short_name, classifier_id, real_value) values($1,$2,$3,$4) RETURNING id")
	} else {
		fmt.Println("некорректный формат значения")
	}
	row := r.db.QueryRow(ctx, query, name, shortName, classifierId, value)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("значение перечисления успешно добавлено")
}

func (r *EnumPositionRepo) DeleteEnumPosition(ctx context.Context, id int) {
	query := fmt.Sprintf("DELETE FROM enum_position WHERE id = $1")
	row, err := r.db.Exec(ctx, query, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rowsAffected := row.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("no rows deleted")
		return
	}
	fmt.Println("значение перечисления успешно удалено")
}

func (r *EnumPositionRepo) ShowEnumPosition(ctx context.Context, id int) {
	var enumPosition models.EnumPosition
	query := fmt.Sprintf("SELECT * FROM enum_position WHERE id = $1")

	row := r.db.QueryRow(ctx, query, id)
	if err := row.Scan(&enumPosition.Id, &enumPosition.Name, &enumPosition.ShortName, &enumPosition.IntegerValue, &enumPosition.RealValue, &enumPosition.StringValue, &enumPosition.ClassifierId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("unit not found")
		}
	}

	fmt.Println(enumPosition.Repr())
}
