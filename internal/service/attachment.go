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

func DeleteAttachmentByUrl(userID int, fileUrl string) error {
	err := d.DeleteAttachmentByUrl(ctx, userID, fileUrl)
	return err
}

func GetAttachmentByUrl(fileUrl string) (*model.Attachment, error) {
	attachment, err := d.GetAttachmentByUrl(ctx, fileUrl)
	return attachment, err
}
