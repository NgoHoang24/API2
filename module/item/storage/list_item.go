package storage

import (
	Profile "Project/module/item/model"
	"context"
)

func (s *sqlStorage) ListProfile(
	ctx context.Context,
	condition map[string]interface{},
	paging *Profile.DataPaging,
) ([]Profile.Profile, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []Profile.Profile

	if err := s.db.Table(Profile.Profile{}.TableName()).
		Where(condition).
		Count(&paging.Total).
		Offset(offset).
		Order("id asc ").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
