package repositories

import (
	"auth/models"
	"database/sql"
)

type AuthRepo interface {
	Registry(model *models.Auth) error
	FindByPhone(phone string) (*models.Auth, error)
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

func (r *authRepo) FindByPhone(phone string) (*models.Auth, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT name, role, phone, password
		FROM auth
		WHERE phone=$1
		LIMIT 1
	`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var auth models.Auth
	err = stmt.QueryRow(phone).Scan(
		&auth.Name, &auth.Role, &auth.Phone, &auth.Password,
	)

	if err != nil {
		return nil, err
	}
	return &auth, tx.Commit()
}
