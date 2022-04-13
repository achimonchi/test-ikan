package repositories

import (
	"auth/models"
	"database/sql"
)

type AuthRepo interface {
	Registry(model *models.Auth) error
}

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) Registry(model *models.Auth) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO auth(id, name, phone, role, password, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		model.ID, model.Name, model.Phone, model.Role, model.Password,
		model.CreatedAt, model.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}
