package storage

import (
	Profile "Project/module/item/model"
	"context"
)

func (s *sqlStorage) CreateProfile(ctx context.Context, data *Profile.Profile) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
func (s *sqlStorage) NewCreateItem(ctx context.Context, data *Profile.Profile) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
