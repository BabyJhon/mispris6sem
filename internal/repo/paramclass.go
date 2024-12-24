package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ParamClassRepo struct {
	db *pgxpool.Pool
}

func NewParamClassRepo(db *pgxpool.Pool) *ParamClassRepo {
	return &ParamClassRepo{
		db: db,
	}
}

func (r *ParamClassRepo) AddParamToClass(ctx context.Context, paramId, prodclassId, minValue, maxValue int) {

	var id int
	query := fmt.Sprintf("INSERT INTO param_class (param_id, prodclass_id, min_value, max_value) values ($1, $2, $3, $4) RETURNING id")
	row := r.db.QueryRow(ctx, query, paramId, prodclassId, minValue, maxValue)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("параметр успешно добавлен классу")
}

func (r *ParamClassRepo) ShowByClass(ctx context.Context, prodClassId int) {
	var paramId int
	query := fmt.Sprintf("SELECT param_id FROM param_class WHERE prodclass_id = $1")

	row := r.db.QueryRow(ctx, query, prodClassId)
	if err := row.Scan(&paramId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("unit not found")
		}
	}

	fmt.Printf("param_id is %d\n", paramId)
}
