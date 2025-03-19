package service

import "github.com/qianqianzyk/AILesson-Planner/internal/model"

func CreateAttachment(attachment *model.Attachment) error {
	err := d.CreateAttachment(ctx, attachment)
	return err
}

func GetAttachmentByUploadID(md5, fileName, uploadID string) (*model.Attachment, error) {
	attachment, err := d.GetAttachmentByUploadID(ctx, md5, fileName, uploadID)
	return attachment, err
}

func UpdateAttachment(attachment *model.Attachment) error {
	err := d.UpdateAttachment(ctx, attachment)
	return err
}
