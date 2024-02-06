package repository

import (
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
)

func (repo *userRepository) SignUp(user model.UserSignUp, avatarURL string) (error, string, string, int, string, int) {

	err, errKey, errMsg, code, isExist := repo.isEmailExist(user.Email)
	if isExist {
		return err, errKey, errMsg, code, getFileInfo("user_sign-up.go"), 0
	}

	stmt, err := repo.db.Preparex(`INSERT INTO users (email, password, name, last_name, avatar) 
        VALUES ($1, $2, $3, $4, $5) RETURNING id`)
	defer stmt.Close()
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("user_sign-up.go"), 0
	}

	var userId int
	if err := stmt.Get(&userId, user.Email, user.Password, user.Name, user.LastName, avatarURL); err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("user_sign-up.go"), 0
	}

	return nil, "", "", http.StatusOK, "", userId
}
