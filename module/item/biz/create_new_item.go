package todobiz

import (
	Profile "Project/module/item/model"
	"context"
)

type CreateProfile interface {
	CreateProfile(ctx context.Context, data *Profile.Profile) error
}

type createBiz struct {
	store CreateProfile
}

func NewCreateProfileBiz(store CreateProfile) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewProfile(ctx context.Context, data *Profile.Profile) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.CreateProfile(ctx, data); err != nil {
		return err
	}

	return nil
}
