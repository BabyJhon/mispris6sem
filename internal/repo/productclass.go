package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductClassRepo struct {
	db *pgxpool.Pool
}

func NewProductClassRepo(db *pgxpool.Pool) *ProductClassRepo {
	return &ProductClassRepo{
		db: db,
	}
}

func (p *ProductClassRepo) AddProductClass(ctx context.Context, className string, unitId, parentId int) {
	var id int
	query := fmt.Sprintf("INSERT INTO prod_class (class_name, unit_id, parent_id) values ($1, $2, $3) RETURNING id")
	row := p.db.QueryRow(ctx, query, className, unitId, parentId)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("класс успешно добавлен")
	return
}

func (p *ProductClassRepo) ChangeParent(ctx context.Context, classId, newParentId int) {
	query := fmt.Sprintf("UPDATE prod_class SET parent_id = $1 WHERE id = $2")
	
	_, err := p.db.Exec(ctx, query, newParentId, classId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("родитель класса успешно обновлен")
}

func (p *ProductClassRepo) SetUnit(ctx context.Context, classId, unitId int) {
	fmt.Println(classId)
	fmt.Println(unitId)
	query := fmt.Sprintf("UPDATE prod_class SET unit_id = $1 WHERE id = $2")
	_, err := p.db.Exec(ctx, query, unitId, classId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("единица измерения успешно обновлена")
}

type Result struct {
	ID          int
	Name        string
	MuID        int
	NameProduct sql.NullString
	Code        sql.NullInt64
}

func (p *ProductClassRepo) PrintChildrenRecursive(ctx context.Context) {
	// Node ID parameter
	nodeID := 1
	// Query to execute
	query := `
WITH RECURSIVE temp1 AS (
SELECT
T1.id,
T1.parent_id,
T1.name,
T1.mu_id,
CAST(T1.name AS VARCHAR(500)) AS path
FROM prod_class T1
WHERE T1.parent_id = $1
UNION ALL
SELECT
T2.id,
T2.parent_id,
T2.name,
T2.mu_id,
CAST(temp1.path || '->' || T2.name AS VARCHAR(500))
FROM prod_class T2
INNER JOIN temp1 ON (temp1.id = T2.parent_id)
)
SELECT
t.id,
t.name,
t.mu_id,
p.name AS name_product,
p.code
FROM temp1 AS t
LEFT JOIN product AS p ON t.id = p.id_class
ORDER BY path;
`

	rows, err := p.db.Query(ctx, query, nodeID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Process the results
	var results []Result
	for rows.Next() {
		var r Result
		err := rows.Scan(&r.ID, &r.Name, &r.MuID, &r.NameProduct, &r.Code)
		if err != nil {
			fmt.Println(err.Error())
		}
		results = append(results, r)
	}
	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		fmt.Print(err.Error())
	}
	// Print the results
	for _, r := range results {
		fmt.Printf("ID: %d, Name: %s, MuID: %d, NameProduct: %s, Code: %d\n",
			r.ID, r.Name, r.MuID, r.NameProduct.String, r.Code.Int64)
	}
}
