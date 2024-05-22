package todobiz

import (
	Profile "Project/module/item/model"
	"context"
)

type DeleteProfileStorage interface {
	FindProfile(
		ctx context.Context,
		condition map[string]interface{},
	) (*Profile.Profile, error)

	DeleteProfile(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteBiz struct {
	store DeleteProfileStorage
}

func NewDeleteProfileBiz(store DeleteProfileStorage) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) DeleteProfile(
	ctx context.Context,
	condition map[string]interface{},
) error {
	// step 1: Find item by conditions
	_, err := biz.store.FindProfile(ctx, condition)

	if err != nil {
		return err
	}

	// Step 2: call to storage to delete item
	if err := biz.store.DeleteProfile(ctx, condition); err != nil {
		return err
	}

	return nil
}
