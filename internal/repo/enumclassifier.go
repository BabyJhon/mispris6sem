package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EnumClassifierRepo struct {
	db *pgxpool.Pool
}

func NewEnumClassifierRepo(db *pgxpool.Pool) *EnumClassifierRepo {
	return &EnumClassifierRepo{
		db: db,
	}
}

func (p *EnumClassifierRepo) AddEnumClassifier(ctx context.Context, name string, parentId int) {
	var id int
	tx, err := p.db.Begin(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	query := fmt.Sprintf("INSERT INTO enum_classifier (name) values ($1) RETURNING id")
	row := tx.QueryRow(ctx, query, name)
	if err := row.Scan(&id); err != nil {
		tx.Rollback(ctx)
		fmt.Println(err.Error())
		return
	}
	query = fmt.Sprintf("UPDATE enum_classifier SET parent_id = $1 WHERE id = $2")
	_, err = tx.Exec(ctx, query, parentId, id)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback(ctx)
		return
	}
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback(ctx)
		return
	}
	fmt.Println("классификатор перечисления успешно добавлен")
}

func (p *EnumClassifierRepo) DeleteEnumClassifier(ctx context.Context, id int) {
	query := fmt.Sprintf("DELETE FROM enum_classifier WHERE id = $1")
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
	fmt.Println("клиссификатор перечисления успешно удален")
}

func (p *EnumClassifierRepo) ChangeParent(ctx context.Context, classId, newParentId int) {
	query := fmt.Sprintf("UPDATE enum_classifier SET parent_id = $1 WHERE id = $2")

	_, err := p.db.Exec(ctx, query, newParentId, classId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("родитель классификатора успешно обновлен успешно обновлен")
}
