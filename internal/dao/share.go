package dao

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
)

func (d *Dao) CreateShareResource(ctx context.Context, post *model.ShareResource) error {
	result := d.orm.WithContext(ctx).Model(&model.ShareResource{}).Create(post)
	return result.Error
}

func (d *Dao) DeleteShareResource(ctx context.Context, postID, userID int) error {
	result := d.orm.WithContext(ctx).
		Where("id = ? AND user_id = ?", postID, userID).
		Delete(&model.ShareResource{})
	return result.Error
}

func (d *Dao) GetShareResourceList(ctx context.Context, resourceType int) ([]model.ShareResource, error) {
	var posts []model.ShareResource
	err := d.orm.WithContext(ctx).Model(&model.ShareResource{}).
		Where("resource_type = ?", resourceType).
		Order("updated_at DESC").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (d *Dao) SearchShareResource(ctx context.Context, resourceType int, keyword string) ([]model.ShareResource, error) {
	var posts []model.ShareResource
	query := "%" + keyword + "%"
	err := d.orm.WithContext(ctx).Where("resource_type = ? AND title LIKE ? OR content LIKE ?", resourceType, query, query).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
