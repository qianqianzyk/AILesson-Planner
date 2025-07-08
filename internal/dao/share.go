package dao

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
)

func (d *Dao) CreateShareResource(ctx context.Context, resource *model.ShareResource) error {
	result := d.orm.WithContext(ctx).Model(&model.ShareResource{}).Create(resource)
	return result.Error
}

func (d *Dao) DeleteShareResource(ctx context.Context, resourceID, userID int) (string, error) {
	var resource model.ShareResource
	err := d.orm.WithContext(ctx).
		Select("cover_img").
		Where("id = ? AND user_id = ?", resourceID, userID).
		First(&resource).Error
	if err != nil {
		return "", err
	}

	result := d.orm.WithContext(ctx).
		Where("id = ? AND user_id = ?", resourceID, userID).
		Delete(&model.ShareResource{})
	return resource.CoverImg, result.Error
}

func (d *Dao) GetShareResourceList(ctx context.Context, resourceType, contentType int) ([]model.ShareResource, error) {
	var resource []model.ShareResource
	query := d.orm.WithContext(ctx).Model(&model.ShareResource{}).Where("resource_type = ?", resourceType)
	if contentType != 0 {
		query = query.Where("content_type = ?", contentType)
	}
	err := query.Order("updated_at DESC").Find(&resource).Error
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (d *Dao) SearchShareResource(ctx context.Context, resourceType, contentType int, keyword string) ([]model.ShareResource, error) {
	var resource []model.ShareResource
	query := "%" + keyword + "%"
	db := d.orm.WithContext(ctx).Where("resource_type = ? AND (title LIKE ? OR content LIKE ?)", resourceType, query, query)
	if contentType != 0 {
		db = db.Where("content_type = ?", contentType)
	}

	err := db.Find(&resource).Error
	if err != nil {
		return nil, err
	}
	return resource, nil
}
