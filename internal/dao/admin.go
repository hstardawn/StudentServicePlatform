package dao

import (
	"StudentServicePlatform/internal/model"
	"context"
)

func (d *Dao) QueryUnhandlePost(ctx context.Context) ([]*model.Post, error) {
	var postList []*model.Post
	err := d.orm.WithContext(ctx).Find(&postList, "status = 0").Error
	return postList, err
}

// func (d *Dao) GetTrash(ctx context.Context)

