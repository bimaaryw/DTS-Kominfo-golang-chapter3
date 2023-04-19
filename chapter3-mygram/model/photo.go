package model

import (
	"time"
)

type Photo struct {
	PhotoID   string `gorm:"primaryKey;type:varchar(255)"`
	Title     string `gorm:"not null;type:varchar(50)"`
	Caption   string `gorm:"type:varchar(255)"`
	PhotoURL  string `gorm:"not null;type:varchar(255)"`
	UserID    string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PhotoCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl" validate:"required"`
}

type PhotoCreateResponse struct {
	PhotoID   string    `json:"photo_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoUpdateRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
}

type PhotoUpdateResponse struct {
	PhotoID string `json:"photo_id"`
}

type PhotoDeleteResponse struct {
	PhotoID string `json:"photo_id"`
}

type PhotoResponse struct {
	PhotoID   string    `json:"photo_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
