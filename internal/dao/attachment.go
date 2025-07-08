package dao

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
)

func (d *Dao) CreateAttachment(ctx context.Context, attachment *model.Attachment) error {
	result := d.orm.WithContext(ctx).Model(&model.Attachment{}).Create(attachment)
	return result.Error
}

func (d *Dao) GetAttachmentByUploadID(ctx context.Context, md5, fileName, uploadID string) (*model.Attachment, error) {
	var attachment *model.Attachment
	result := d.orm.WithContext(ctx).Model(&model.Attachment{}).
		Where("md5 = ? AND file_name = ? AND upload_id = ?", md5, fileName, uploadID).First(&attachment)
	if result.Error != nil {
		return nil, result.Error
	}
	return attachment, nil
}

func (d *Dao) UpdateAttachment(ctx context.Context, attachment *model.Attachment) error {
	result := d.orm.WithContext(ctx).Model(&model.Attachment{}).Save(attachment)
	return result.Error
}

func (d *Dao) DeleteAttachmentByUrl(ctx context.Context, userID int, fileUrl string) error {
	result := d.orm.WithContext(ctx).
		Where("user_id = ? AND file_url = ?", userID, fileUrl).
		Delete(&model.Attachment{})
	return result.Error
}

func (d *Dao) GetAttachmentByUrl(ctx context.Context, fileUrl string) (*model.Attachment, error) {
	var attachment *model.Attachment
	result := d.orm.WithContext(ctx).Model(&model.Attachment{}).
		Where("file_url = ?", fileUrl).
		First(&attachment)
	return attachment, result.Error
}
