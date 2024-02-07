package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.User
}

func newUserService(repo repository.User) *userService {
	return &userService{repo: repo}
}

// ! ================================== image handle ========================================

func (service *userService) imageHandler(file *multipart.FileHeader, email string) (error, string, string, int, string, string) {
	if !service.imageValidation(file) {
		return errors.New("Incorrect file type"), "Bad Request", "Incorrect file type", http.StatusBadRequest, "", ""
	}

	err, avatarPath := service.saveImage(file, email)
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusBadRequest, "", ""
	}

	avatarURL := service.constructURL(avatarPath)
	return nil, "", "", 200, avatarURL, avatarPath
}

func (service *userService) imageValidation(file *multipart.FileHeader) bool {
	extension := strings.ToLower(filepath.Ext(file.Filename))
	if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
		return false
	}

	return true
}

func (service *userService) saveImage(file *multipart.FileHeader, email string) (error, string) {
	folderPath := "app/pkg/storage/avatar_image/" + email
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return err, ""
	}

	filePath := filepath.Join(folderPath, file.Filename)
	return nil, filePath
}

func (service *userService) constructURL(filePath string) string {
	filePath = strings.TrimPrefix(filePath, "app/pkg/storage")
	apiPort := os.Getenv("API_PORT")

	return "http://localhost:" + apiPort + filePath
}

// ! ================================== validation ========================================

func (service *userService) validateSignUp(email, password, name, lastName string) string {
	if errMsg := service.isValidPassword(password); errMsg != "" {
		return errMsg
	}

	if errMsg := service.isValidEmail(email); errMsg != "" {
		return errMsg
	}

	if errMsg := service.isValidName(name); errMsg != "" {
		return errMsg
	}

	if errMsg := service.isValidLastName(lastName); errMsg != "" {
		return errMsg
	}

	return ""
}

func (service *userService) validateSignIn(email, password string) string {
	if errMsg := service.isValidPassword(password); errMsg != "" {
		return errMsg
	}

	if errMsg := service.isValidEmail(email); errMsg != "" {
		return errMsg
	}

	return ""
}

func (service *userService) isValidPassword(password string) string {
	if len(password) < 8 || len(password) > 64 {
		return "Password validation is not allowed"
	}

	return ""
}

func (service *userService) isValidEmail(email string) string {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)

	if !match {
		return "Invalid email format"
	}

	return ""
}

func (service *userService) isValidName(name string) string {
	if len(name) == 0 {
		return "Name cannot be empty"
	}
	return ""
}

func (service *userService) isValidLastName(lastName string) string {
	if len(lastName) == 0 {
		return "Last name cannot be empty"
	}
	return ""
}

// ! ================================== password ========================================

func (service *userService) passwordHasher(password string) string {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(password+os.Getenv("HASH_SALT")), 10)
	return string(newPassword)
}

func (service *userService) comparePasswords(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+os.Getenv("HASH_SALT")))
}

// ! ================================== tokens creation ========================================

func (service *userService) tokensHandler(userId int) (error, model.Tokens) {
	var err error
	var tokens model.Tokens

	tokens.AccessToken, err = CreteAccessToken(userId)
	if err != nil {
		return err, tokens
	}

	tokens.RefreshToken, err = CreateRefreshToken(userId)
	if err != nil {
		return err, tokens
	}

	return nil, tokens
}

func CreteAccessToken(userId int) (string, error) {
	claims := jwt.MapClaims{"sub": userId, "exp": time.Now().Add(time.Minute * 1).Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))
}

func CreateRefreshToken(userId int) (string, error) {
	claims := jwt.MapClaims{"sub": userId, "exp": time.Now().Add(time.Hour * 24 * 30).Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("REFRESH_SECRET_KEY")))
}

func IsTokenValid(tokenString string, isRefreshToken bool) (error, string, string, int, string, int) {
	secretKey := os.Getenv("ACCESS_SECRET_KEY")
	if isRefreshToken {
		secretKey = os.Getenv("REFRESH_SECRET_KEY")
	}

	if tokenString == "" {
		return errors.New("Token is empty"), "Unauthorized", "Token is empty", http.StatusUnauthorized, getFileInfo("user.go"), 0
	}

	tokenSlice := strings.Split(tokenString, " ")
	if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
		return errors.New("Invalid authorization header format"), "Unauthorized", "Invalid authorization header format", http.StatusUnauthorized, getFileInfo("user.go"), 0
	}

	fmt.Println(tokenString, tokenSlice[0], tokenSlice[1])
	token, err := jwt.Parse(tokenSlice[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return errors.New("Invalid token"), "Unauthorized", "Token is malformed", http.StatusUnauthorized, getFileInfo("user.go"), 0
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return errors.New("Invalid token"), "Unauthorized", "Token has expired", http.StatusUnauthorized, getFileInfo("user.go"), 0
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return errors.New("Invalid token"), "Unauthorized", "Token not yet valid", http.StatusUnauthorized, getFileInfo("user.go"), 0
			} else {
				return errors.New("Invalid token"), "Unauthorized", "Couldn't handle this token", http.StatusUnauthorized, getFileInfo("user.go"), 0
			}
		}
		return errors.New("Invalid token"), "Unauthorized", "Couldn't handle this token", http.StatusUnauthorized, getFileInfo("user.go"), 0
	}

	if !token.Valid {
		return errors.New("Invalid token"), "Unauthorized", "Invalid token", http.StatusUnauthorized, getFileInfo("user.go"), 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("Invalid token claims"), "Unauthorized", "Invalid token claims", http.StatusUnauthorized, getFileInfo("user.go"), 0
	}

	userIdFloat, ok := claims["sub"].(float64)
	if !ok {
		return errors.New("Invalid userId in token"), "Unauthorized", "Invalid userId in token", http.StatusUnauthorized, getFileInfo("user.go"), 0
	}

	userId := int(userIdFloat)
	return nil, "", "", http.StatusOK, "", userId
}
