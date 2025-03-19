package service

import (
	"bytes"
	"fmt"
	"github.com/minio/minio-go/v7"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"github.com/dustin/go-humanize"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

const ImageLimit = humanize.MByte * 5

func GenerateObjectKey(uploadType string, fileExt string) string {
	extFolder := strings.TrimPrefix(fileExt, ".")
	return fmt.Sprintf("%s/%d/%s/%s%s",
		uploadType,
		time.Now().Year(),
		extFolder,
		uuid.NewV1().String(),
		fileExt,
	)
}

func ConvertToWebP(reader io.Reader) (io.Reader, int64, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, 0, err
	}

	var buf bytes.Buffer
	err = webp.Encode(&buf, img, &webp.Options{Quality: 100})
	if err != nil {
		return nil, 0, err
	}
	return bytes.NewReader(buf.Bytes()), int64(buf.Len()), nil
}

func PutObject(objectKey string, reader io.Reader, size int64, contentType string) (string, error) {
	url, err := d.PutObject(objectKey, reader, size, contentType)
	return url, err
}

func DeleteObjectByUrlAsync(url string) {
	objectKey, ok := d.GetObjectKeyFromUrl(url)
	if ok {
		go func(objectKey string) {
			err := d.DeleteObject(objectKey)
			if err != nil {
				zap.L().Error("Minio 删除对象错误", zap.String("objectKey", objectKey), zap.Error(err))
			}
		}(objectKey)
	}
}

func CheckFileMD5(md5, fileName string) (status int, uploadID string, uploadedParts []minio.CompletePart, err error) {
	status, uploadID, uploadedParts, err = d.CheckFileMD5(ctx, md5, fileName)
	return status, uploadID, uploadedParts, err
}

func InitMultipartUpload(md5, fileName string, contentType string) (uploadID string, err error) {
	objectKey := "disk/" + md5 + "_" + fileName
	uploadID, err = d.InitMultipartUpload(ctx, objectKey, contentType)
	return uploadID, err
}

func CompleteMultipartUpload(md5, fileName string, uploadID string, contentType string) (string, int64, error) {
	objectKey := "disk/" + md5 + "_" + fileName
	fileUrl, size, err := d.CompleteMultipartUpload(ctx, objectKey, uploadID, contentType)
	return fileUrl, size, err
}
