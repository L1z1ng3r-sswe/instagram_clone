package model

type Tokens struct {
	AccessToken  string `form:"access_token"`
	RefreshToken string `form:"refresh_token"`
}

// For sign-up
type UserSignUp struct {
	Email    string `form:"email" db:"email"`
	Password string `form:"password" db:"password"`
	Name     string `form:"name" db:"name"`
	LastName string `form:"last_name" db:"last_name"`
}

// For sign-in
type UserSignIn struct {
	Email    string `form:"email" db:"email"`
	Password string `form:"password" db:"password"`
}

// For get
type User struct {
	Id       int    `form:"id" db:"id"`
	Email    string `form:"email" db:"email"`
	Name     string `form:"name" db:"name"`
	LastName string `form:"last_name" db:"last_name"`
	Avatar   string `form:"avatar" db:"avatar"`
}

// For password change and update
type UserUpdate struct {
	Id       int    `form:"id" db:"id"`
	Email    string `form:"email" db:"email"`
	Name     string `form:"name" db:"name"`
	LastName string `form:"last_name" db:"last_name"`
}
