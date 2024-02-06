package repository

import (
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
)

func (repo *userRepository) SignIn(user model.UserSignIn) (error, string, string, int, string, string, int) {
	err, errKey, errMsg, code, isExist := repo.isEmailExist(user.Email)
	if !isExist {
		return err, errKey, errMsg, code, getFileInfo("user_sign-in.go"), "", 0
	}

	stmt, err := repo.db.Preparex(`SELECT id, password FROM users WHERE email = $1 LIMIT 1`)
	defer stmt.Close()
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("user_sign-in.go"), "", 0
	}

	type UserDB struct {
		Id       int    `db:"id"`
		Password string `db:"password"`
	}

	var userDB UserDB
	if err := stmt.Get(&userDB, user.Email); err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("user_sign-in.go"), "", 0
	}

	return nil, "", "", http.StatusOK, "", userDB.Password, userDB.Id
}
