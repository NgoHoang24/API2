package storage

import (
	Profile "Project/module/item/model"
	"context"
)

func (s *sqlStorage) UpdateProfile(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *Profile.Profile,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
