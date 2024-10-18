package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const TableName = "visits"

type VisitRepository struct {
	db *sqlx.DB
}

func NewVisitRepository(db *sqlx.DB) *VisitRepository {
	return &VisitRepository{
		db: db,
	}
}

func (r *VisitRepository) GetCount() (int, error) {
	q := fmt.Sprintf(`SELECT COUNT(*) FROM %s`, TableName)
	var count int

	if err := r.db.Get(&count, q); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *VisitRepository) Create(visit Visit) (int, error) {
	q := fmt.Sprintf("INSERT INTO %s (`time`, `ip`, `city`) VALUES (:time, :ip, :city)", TableName)

	_, err := r.db.NamedExec(q, visit)
	if err != nil {
		return 0, err
	}

	return r.GetCount()
}
