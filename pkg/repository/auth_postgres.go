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
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash, referralcode) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.UserName, user.Email, user.Password, user.ReferralCode)
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
func (r *AuthPostgres) CreateCounter() (int, error) {
	var count int

	query := fmt.Sprintf("SELECT count(*) FROM %s", usersTable)

	row := r.db.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
