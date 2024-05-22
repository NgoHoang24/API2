package todobiz

import (
	Profile "Project/module/item/model"
	"context"
)

type FindProfileStorage interface {
	FindProfile(
		ctx context.Context,
		condition map[string]interface{},
	) (*Profile.Profile, error)
}

type findBiz struct {
	store FindProfileStorage
}

func NewFindProfileBiz(store FindProfileStorage) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) FindAProfile(ctx context.Context, condition map[string]interface{}) (*Profile.Profile, error) {
	itemData, err := biz.store.FindProfile(ctx, condition)

	if err != nil {
		return nil, err
	}

	return itemData, nil
}
