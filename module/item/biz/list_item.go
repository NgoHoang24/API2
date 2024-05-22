package todobiz

import (
	Profile "Project/module/item/model"
	"context"
)

type ListProfileStorage interface {
	ListProfile(
		ctx context.Context,
		condition map[string]interface{},
		paging *Profile.DataPaging,
	) ([]Profile.Profile, error)
}

type listProfileBiz struct {
	store ListProfileStorage
}

func NewListProfileBiz(store ListProfileStorage) *listProfileBiz {
	return &listProfileBiz{store: store}
}

func (biz *listProfileBiz) ListProfile(ctx context.Context,
	condition map[string]interface{},
	paging *Profile.DataPaging,
) ([]Profile.Profile, error) {
	result, err := biz.store.ListProfile(ctx, condition, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
