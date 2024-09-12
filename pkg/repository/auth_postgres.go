package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	restJournal "rest_journal"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user restJournal.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, username, email, password, role) values ($1, $2, $3, $4, $5) RETURNING id)`)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
