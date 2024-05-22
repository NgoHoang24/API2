package storage

import (
	Profile "Project/module/item/model"
	"context"
)

func (s *sqlStorage) FindProfile(
	ctx context.Context,
	condition map[string]interface{},
) (*Profile.Profile, error) {
	var itemData Profile.Profile

	if err := s.db.Where(condition).First(&itemData).Error; err != nil {
		return nil, err // other errors
	}

	return &itemData, nil
}
