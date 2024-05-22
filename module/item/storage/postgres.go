package storage

import (
	Profile "Project/module/item/model"
	"context"
	"gorm.io/gorm"
)

type sqlStorage struct {
	db *gorm.DB
}

func (s *sqlStorage) ListProfileStorage(ctx context.Context, condition map[string]interface{}, paging *Profile.DataPaging) ([]Profile.Profile, error) {
	panic("implement me")
}

func NewSQLStorage(db *gorm.DB) *sqlStorage {
	return &sqlStorage{db: db}
}
