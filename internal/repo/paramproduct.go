package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ParamProductRepo struct {
	db *pgxpool.Pool
}

func NewParamProductRepo(db *pgxpool.Pool) *ParamProductRepo {
	return &ParamProductRepo{
		db: db,
	}
}

func (p *ParamProductRepo) AddParamToProduct(ctx context.Context, value, paramClassId, productId int) {
	var minValue, maxValue int
	query := fmt.Sprintf("SELECT min_value, max_value FROM param_class WHERE id = $1;")
	row := p.db.QueryRow(ctx, query, paramClassId)
	if err := row.Scan(&minValue, &maxValue); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("unit not found")
		}
	}

	if value < minValue || value > maxValue {
		fmt.Println("значение не попадает в диапазон")
		return
	}

	var id int
	query = fmt.Sprintf("INSERT INTO param_product (value, product_id, param_class_id) values ($1, $2, $3) RETURNING id")

	row = p.db.QueryRow(ctx, query, value, productId, paramClassId)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("параметр успешно добавлен продукту")
	return
}

func (p *ParamProductRepo) Edit(ctx context.Context, productId, paramClassId, value int) {
	var minValue, maxValue int
	query := fmt.Sprintf("SELECT min_value, max_value FROM param_class WHERE id = $1;")
	row := p.db.QueryRow(ctx, query, paramClassId)
	if err := row.Scan(&minValue, &maxValue); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("unit not found")
		}
	}

	query = fmt.Sprintf("UPDATE param_product SET value = $1 WHERE param_class_id = $2")
	_, err := p.db.Exec(ctx, query, value, paramClassId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("параметр успешно обновлен")
}

func (p *ParamProductRepo) ShowAllByProduct(ctx context.Context, productId int) {
	var id, value, pId, paramClassId int
	query := fmt.Sprintf("SELECT * FROM param_product WHERE product_id = $1")

	rows, err := p.db.Query(ctx, query, productId)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer rows.Close()

	for rows.Next() {
		//var book entity.Book
		if err := rows.Scan(&id, &value, &pId, &paramClassId); err != nil {
			fmt.Println("error")
			return
		}
		fmt.Printf("id=%d, value=%d, product_id=%d, param_class_id=%d", id, value, pId, paramClassId)
	}
}
