package model

import "time"

type Response struct {
	AdminID        int       `json:"admin_id"`
	Response       string    `json:"response"`
    ResponseRating int       `json:"response_rating"`
	// UserID         int       `json:"user_id"`
	PostID         int       `json:"post_id"`
	// PostType       int       `json:"post_type"`//后续看管理员擅长什么领域? 便于超级管理员管理?
	// Title          string    `json:"title"`
	// Content        string    `json:"content"`
	// Status         int       `json:"status"`
	CreateAt       time.Time `json:"response_time"`
	// UpdatedAt      time.Time `json:"updated_response"`
}
