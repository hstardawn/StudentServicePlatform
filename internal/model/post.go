package model

import "time"

type Post struct {
	Id             int       `json:"post_id"`
	UserId         int       `json:"user_id"`
	IsAnonymous    int       `json:"is_anonymous"`
	IsUrgent       int       `json:"is_urgent"`
	PostType       int       `json:"post_type"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Status         int       `json:"status"`
	Response       string    `json:"response"`
	ResponseRating int       `json:"response_rating"`
	CreateAt       time.Time `json:"time"`
	UpdataAt       time.Time `json:"updated_at"`
}