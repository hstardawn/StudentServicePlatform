package service

import (
	"path"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func IsImage(filename string) bool {
	ext := path.Ext(filename)
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp":
		return true
	default:
		return false
	}
}

func GetFileType(filename string) string {
	return filepath.Ext(filename)
}

func GetUUID() string {
	return uuid.NewV1().String()
}

func StorePostPicture(ID int,filename string) error {
	return d.StorePostPicture(ctx,ID,filename)
}

func StoreUserPicture(ID int,filename string) error {
	return d.StoreUserPicture(ctx,ID,filename)
}