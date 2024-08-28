package repository

import (
	"database/sql"
	"errors"

	"github.com/go-kit/log"
	"github.com/mauricioww/goauth/app/models"
)

type Repository struct {
	db     *sql.DB
	logger log.Logger
}

type Repositorier interface {
	CreateUser(firstName string, lastName string, password string, admin bool) (bool, error)
	GetUser(email string) (models.User, error)
}

func NewRepository(db *sql.DB, l log.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: l,
	}
}

func (r *Repository) CreateUser(firstName string, lastName string, password string, admin bool) (bool, error) {
	_, err := r.db.Exec(CreateUserQ, firstName, lastName, password, admin)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) GetUser(email string) (models.User, error) {
	var u models.User

	err := r.db.QueryRow(GetUserQ, email).Scan(u.FirstName, u.LastName, u.Email, u.Admin, u.Active)

	if err == sql.ErrNoRows {
		return models.User{}, errors.New("not found")
	}

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
