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
