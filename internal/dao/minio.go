package dao

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"gorm.io/gorm"
	"io"
	"strings"
)

// PutObject 用于上传对象
func (d *Dao) PutObject(objectKey string, reader io.Reader, size int64, contentType string) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	_, err := (*d.mn).Client.PutObject(context.Background(), (*d.mn).Bucket, objectKey, reader, size, "", "", opts)
	if err != nil {
		return "", err
	}
	return (*d.mn).Domain + (*d.mn).Bucket + "/" + objectKey, nil
}

// GetObjectKeyFromUrl 从 Url 中提取 ObjectKey
// 若该 Url 不是来自此 Minio, 则 ok 为 false
func (d *Dao) GetObjectKeyFromUrl(fullUrl string) (objectKey string, ok bool) {
	objectKey = strings.TrimPrefix(fullUrl, (*d.mn).Domain+(*d.mn).Bucket+"/")
	if objectKey == fullUrl {
		return "", false
	}
	return objectKey, true
}

// DeleteObject 用于删除相应对象
func (d *Dao) DeleteObject(objectKey string) error {
	err := (*d.mn).Client.RemoveObject(
		context.Background(),
		(*d.mn).Bucket,
		objectKey,
		minio.RemoveObjectOptions{ForceDelete: true},
	)
	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}
	return nil
}

func (d *Dao) InitMultipartUpload(ctx context.Context, objectKey string, contentType string) (uploadID string, err error) {
	uploadID, err = (*d.mn).Client.NewMultipartUpload(ctx, (*d.mn).Bucket, objectKey, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("初始化上传失败: %v", err)
	}
	return uploadID, nil
}

func (d *Dao) CheckFileMD5(ctx context.Context, md5, fileName string) (status int, uploadID string, uploadedParts []minio.CompletePart, err error) {
	var attachment model.Attachment
	err = d.orm.WithContext(ctx).Where("md5 = ? AND file_name = ?", md5, fileName).First(&attachment).Error
	if err == nil {
		return 1, attachment.FileUrl, nil, nil
	} else if err != gorm.ErrRecordNotFound {
		return 0, "", nil, err
	}

	objectKey := "disk/" + md5 + "_" + fileName
	parts, err := d.getUploadedParts(ctx, objectKey, attachment.UploadID)
	if err != nil {
		return 0, "", nil, err
	}
	if uploadID != "" {
		return 2, attachment.UploadID, parts, nil
	}
	return 3, "", nil, nil
}

func (d *Dao) getUploadedParts(ctx context.Context, objectKey, uploadID string) ([]minio.CompletePart, error) {
	// 定义存储已上传的分片
	var parts []minio.CompletePart
	// 记录分页的起始 partNumber
	partNumberMarker := 0
	for {
		// 调用 ListObjectParts 获取分片信息
		resp, err := (*d.mn).Client.ListObjectParts(ctx, (*d.mn).Bucket, objectKey, uploadID, partNumberMarker, 1000)
		if err != nil {
			return nil, fmt.Errorf("获取已上传分块失败: %v", err)
		}
		// 解析返回的 Part 信息
		for _, part := range resp.ObjectParts {
			parts = append(parts, minio.CompletePart{
				PartNumber: part.PartNumber,
				ETag:       part.ETag,
			})
		}
		// 如果已经获取所有分片，跳出循环
		if !resp.IsTruncated {
			break
		}
		// 继续请求下一个分页
		partNumberMarker = resp.NextPartNumberMarker
	}
	return parts, nil
}

// CompleteMultipartUpload 用于合并文件
func (d *Dao) CompleteMultipartUpload(ctx context.Context, objectKey, uploadID string, contentType string) (string, int64, error) {
	uploadedParts, err := d.getUploadedParts(ctx, objectKey, uploadID)
	if err != nil {
		_ = (*d.mn).Client.AbortMultipartUpload(ctx, (*d.mn).Bucket, objectKey, uploadID)
		return "", 0, err
	}

	_, err = (*d.mn).Client.CompleteMultipartUpload(ctx, (*d.mn).Bucket, objectKey, uploadID, uploadedParts, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		_ = (*d.mn).Client.AbortMultipartUpload(ctx, (*d.mn).Bucket, objectKey, uploadID)
		return "", 0, err
	}
	fileUrl := fmt.Sprintf("%s%s/%s", (*d.mn).Domain, (*d.mn).Bucket, objectKey)

	err = (*d.mn).Client.AbortMultipartUpload(ctx, (*d.mn).Bucket, objectKey, uploadID)
	if err != nil {
		return "", 0, err
	}

	stat, err := (*d.mn).Client.StatObject(ctx, d.mn.Bucket, objectKey, minio.StatObjectOptions{})
	if err != nil {
		return "", 0, err
	}

	return fileUrl, stat.Size, nil
}
