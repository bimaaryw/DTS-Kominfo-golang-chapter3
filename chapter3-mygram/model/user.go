package model

import (
	"time"
)

type User struct {
	UserID      string `gorm:"primaryKey;type:varchar(255)"`
	Username    string `gorm:"not null;uniqueIndex;type:varchar(50)"`
	Email       string `gorm:"not null;uniqueIndex;type:varchar(50)"`
	Password    string `gorm:"not null;type:varchar(255)"`
	Age         int64  `gorm:"not null;type:integer"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Photos      []Photo
	SocialMedia []SocialMedia
	Comment     []Comment
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int64  `json:"age" validate:"required,gt=8"`
}

type UserRegisterResponse struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int64     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	UserID      string        `json:"user_id"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Age         int64         `json:"age"`
	Photos      []Photo       `json:"photos"`
	SocialMedia []SocialMedia `json:"social media"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
