package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ParamRepo struct {
	db *pgxpool.Pool
}

func NewParamRepo(db *pgxpool.Pool) *ParamRepo {
	return &ParamRepo{
		db: db,
	}
}

func (p *ParamRepo) AddParam(ctx context.Context, name, shortName string, unitId, enumClassifierId int) {
	var id int
	query := fmt.Sprintf("INSERT INTO param (name, short_name, unit_id, enum_classifier_id) values ($1, $2, $3, $4) RETURNING id")

	row := p.db.QueryRow(ctx, query, name, shortName, unitId, enumClassifierId)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("параметр успешно добавлен")
	return
}
