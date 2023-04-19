package model

import (
	"time"
)

type Comment struct {
	CommentID string `gorm:"primaryKey;type:varchar(255)"`
	Message   string `gorm:"not null;type:varchar(255)"`
	UserID    string
	PhotoID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentCreateRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentUpdateResponse struct {
	CommentID string `json:"comment_id"`
}

type CommentDeleteResponse struct {
	CommentID string `json:"comment_id"`
}

type CommentResponse struct {
	CommentID string    `json:"comment_id"`
	Message   string    `json:"message"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
