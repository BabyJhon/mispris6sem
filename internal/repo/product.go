package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/BabyJhon/mispris1-2/internal/models"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) AddProduct(ctx context.Context, productName string, productClassId, enumClassId int) {

	var id int
	query := fmt.Sprintf("INSERT INTO product (product_name, prod_class_id, enum_classifier_id) values ($1, $2, $3) RETURNING id_product")

	row := p.db.QueryRow(ctx, query, productName, productClassId, enumClassId)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("продукт успешно добавлен")
	return
}

func (p *ProductRepo) DeleteProduct(ctx context.Context, id int) {
	query := fmt.Sprintf("DELETE FROM product WHERE id_product = $1")
	row, err := p.db.Exec(ctx, query, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rowsAffected := row.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("no rows deleted")
		return
	}
	fmt.Println("продукт успешно удален")
}

func (p *ProductRepo) ShowProduct(ctx context.Context, id int) {
	var product models.Product
	query := fmt.Sprintf("SELECT * FROM product WHERE id_product = $1")

	row := p.db.QueryRow(ctx, query, id)
	if err := row.Scan(&product.ID, &product.ProductName, &product.ProductClassId, &product.EnumClassifierId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("unit not found")
		}
	}

	fmt.Println(product.Repr())
}

func (p *ProductRepo) ShowByProdClass(ctx context.Context, classId int) {
	//var products []models.Product
	query := fmt.Sprintf("SELECT * FROM product WHERE prod_class_id = $1")

	rows, err := p.db.Query(ctx, query, classId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.ProductName, &product.ProductClassId, &product.EnumClassifierId); err != nil {
			return
		}
		fmt.Println(product.Repr())
	}
}

func (p *ProductRepo) CheckClass(ctx context.Context, productId, classId int) {
	currentId := productId
	query := "SELECT prod_class_id FROM product WHERE id_product = $1"
	var currentParentId int
	row := p.db.QueryRow(ctx, query, currentId)
	if err := row.Scan(&currentParentId); err != nil {
		fmt.Println(err.Error())
	}

	for currentId != currentParentId {
		query := "SELECT id, parent_id FROM prod_class WHERE id = $1"
		row := p.db.QueryRow(ctx, query, currentParentId)
		if err := row.Scan(&currentId, &currentParentId); err != nil {
			fmt.Println(err.Error())
		}

		if classId == currentParentId {
			fmt.Println("изделие принадлежит заданному классу")
			return
		}
	}
	fmt.Println("издлие не принадлежит заданному классу")
}
