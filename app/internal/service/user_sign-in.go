package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
)

func (service *userService) SignIn(user model.UserSignIn) (error, string, string, int, string, int, model.Tokens) {
	if errMsg := service.validateSignIn(user.Email, user.Password); errMsg != "" {
		return errors.New("Validation failed"), "Bad Request", errMsg, http.StatusBadRequest, getFileInfo("user_sign-in.go"), 0, model.Tokens{}
	}

	err, errKey, errMsg, code, fileInfo, hashedPassword, userId := service.repo.SignIn(user)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo, 0, model.Tokens{}
	}

	fmt.Println("hashedPassword: "+hashedPassword, "password: ", user.Password)
	if err := service.comparePasswords(user.Password, hashedPassword); err != nil {
		return err, "Bad Request", err.Error(), http.StatusBadRequest, getFileInfo("user_sign-in.go"), 0, model.Tokens{}
	}

	err, tokens := service.tokensHandler(userId)
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("user_sign-in.go"), 0, model.Tokens{}
	}

	return nil, "", "", http.StatusOK, "", userId, tokens
}
