package repository

import (
	registration "TBCC_RegistrationPJ"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user registration.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthPostgres) GetUser(Email, password string) (registration.User, error) {
	var user registration.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE Email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, Email, password)
	return user, err
}
