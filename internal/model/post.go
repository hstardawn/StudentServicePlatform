package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID             int       `json:"post_id"`
	UserID         int       `json:"user_id"`
	Name           string    `json:"name"`
	IsAnonymous    int       `json:"is_anonymous"`
	IsUrgent       int       `json:"is_urgent"`
	PostType       int       `json:"post_type"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
    Pictures       string    `json:"pictures"`
	Status         int       `json:"status"`
	AdminID 	   int       `json:"admin_id"`
	// Response       string    `json:"response"`
	// ResponseRating int       `json:"response_rating"`
	CreateAt       time.Time `json:"post_time"`
	UpdatedAt      time.Time `json:"updated_post"`
	// ResponseAt     time.Time `json:"response_time"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
}
