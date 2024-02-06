package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) isEmailExist(email string) (error, string, string, int, bool) {
	stmt, err := repo.db.Preparex(`SELECT id FROM users WHERE email=$1 LIMIT 1`)
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, true
	}
	defer stmt.Close()

	var userIdCheck int
	if err = stmt.Get(&userIdCheck, email); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Not Found"), "Not Found", "User account does not exist", http.StatusNotFound, false
		} else {
			return err, "Not Found", "User account does not exist", http.StatusNotFound, false
		}
	}

	return errors.New("Bad Request"), "Bad Request", "Email is already exist", http.StatusBadRequest, true
}
