package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/BabyJhon/mispris1-2/internal/models"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitRepo struct {
	db *pgxpool.Pool
}

func NewUnitRepo(db *pgxpool.Pool) *UnitRepo {
	return &UnitRepo{db: db}
}

func (u *UnitRepo) AddUnit(ctx context.Context, unitName, shortName string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO unit (unit_name, short_name) values ($1, $2) RETURNING id")

	row := u.db.QueryRow(ctx, query, unitName, shortName)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	fmt.Println("единица измерения успешно добавлена")
	return id, nil
}

func (u *UnitRepo) DeleteUnit(ctx context.Context, id int) {
	query := fmt.Sprintf("DELETE FROM unit WHERE id = $1")
	row, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return
	}

	rowsAffected := row.RowsAffected()
	if rowsAffected == 0 {
		return
	}
	fmt.Println("единица измерения успешно удалена")
}

func (u *UnitRepo) ShowUnit(ctx context.Context, id int) {
	var unit models.Unit
	query := fmt.Sprintf("SELECT * FROM unit WHERE id = $1")

	row := u.db.QueryRow(ctx, query, id)
	if err := row.Scan(&unit.ID, &unit.Unit_name, &unit.Short_name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("unit not found")
		}
	}

	fmt.Println(unit.Repr())
}

func (u *UnitRepo) UpdateUnit(ctx context.Context, id int, name, shortName string) {
	query := fmt.Sprintf("UPDATE unit SET unit_name = $1, short_name = $2 WHERE id = $3")
	// fmt.Println(classId)
	// fmt.Println(newParentId)
	_, err := u.db.Exec(ctx, query, name, shortName, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("единица измерения успешно обновлена")
}