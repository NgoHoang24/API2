package todobiz

import (
	Profile "Project/module/item/model"
	"context"
)

type UpdateProfileStorage interface {
	FindProfile(
		ctx context.Context,
		condition map[string]interface{},
	) (*Profile.Profile, error)

	UpdateProfile(
		ctx context.Context,
		condition map[string]interface{},
		dataUpdate *Profile.Profile,
	) error
}

type updateBiz struct {
	store UpdateProfileStorage
}

func NewUpdateProfileBiz(store UpdateProfileStorage) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) UpdateProfile(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *Profile.Profile,
) error {
	// step 1: Find item by conditions
	oldItem, err := biz.store.FindProfile(ctx, condition)

	if err != nil {
		return err
	}

	// just a demo in case we dont allow update Finished item
	if oldItem.Email == "Finished" {
	}

	// Step 2: call to storage to update item
	if err := biz.store.UpdateProfile(ctx, condition, dataUpdate); err != nil {
		return err
	}

	return nil
}
