package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)


type User struct{
	ID string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	IsAdmin bool `json:"is_admin" bson:"is_admin"`
	IsActive bool `json:"is_active" bson:"is_active"`
	VerifyToken string `json:"verify_token" bson:"verify_token"`
	ExpirationDatetoken time.Time `json:"expiration_date_token" bson:"expiration_date_token"`
	
}

type UserRepo interface{
	 CreateUser(user *User) error
	GetUser(id string) (*User, error)	
	GetUsers() ([]*User, error)
	GetUserByEmail(email string) (User, error)
	MakeAcitiveUser(user User) (User, error)
	UpdateUsertoken(user *User) error
	UpdatePassword(user *User) error
	DeleteUser(id string) error
}
type UserUsecase interface{
	AddeUser(user *User) (*User, error)
	GetUser(id string) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}
type ForgetPassword struct{
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	VerifyToken string `json:"verify_token" bson:"verify_token"`
	ConfrimPassword string `json:"confirm_password" bson:"confirm_password"`
}
type Claims struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	jwt.StandardClaims
}