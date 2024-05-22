package storage

import (
	Profile "Project/module/item/model"
	"context"
)

func (s *sqlStorage) DeleteProfile(
	ctx context.Context,
	condition map[string]interface{},
) error {

	if err := s.db.
		Table(Profile.Profile{}.TableName()).
		Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
