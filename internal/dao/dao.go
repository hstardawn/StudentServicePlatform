package dao

import (
	"StudentServicePlatform/internal/model"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Dao struct {
	orm *gorm.DB
}

func New(orm *gorm.DB) *Dao {
	return &Dao{orm: orm}
}

type Daos interface {
	//user
	GetUserByUserID(ctx context.Context, userID int)

	// student
	GetPostByPostID(ctx context.Context, postID int)
	SavePost(ctx context.Context, post model.Post)

	//admin
	QueryUnhandlePost(ctx context.Context)
	ReceivePost()
}