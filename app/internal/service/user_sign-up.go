package service

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
)

func (service *userService) SignUp(user model.UserSignUp, avatar *multipart.FileHeader) (error, string, string, int, string, string, model.Tokens, int) {
	// validate
	if errMsg := service.validateSignUp(user.Email, user.Password, user.Name, user.LastName); errMsg != "" {
		return errors.New("Validation failed"), "Bad Request", errMsg, http.StatusBadRequest, getFileInfo("user_sign-up.go"), "", model.Tokens{}, 0
	}

	// handle image
	err, errKey, errMsg, code, avatarURL, avatarPath := service.imageHandler(avatar, user.Email)
	if err != nil {
		return err, errKey, errMsg, code, getFileInfo("user_sing-up.go"), "", model.Tokens{}, 0
	}

	// db
	user.Password = service.passwordHasher(user.Password)
	err, errKey, errMsg, code, fileInfo, userId := service.repo.SignUp(user, avatarURL)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo, "", model.Tokens{}, 0
	}

	// tokens creation
	err, tokens := service.tokensHandler(userId)
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("user_sign-up.go"), "", model.Tokens{}, 0
	}

	return nil, "", "", http.StatusOK, "", avatarPath, tokens, userId
}
