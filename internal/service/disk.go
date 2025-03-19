package service

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/big"
	"regexp"
	"strings"
	"time"
)

func CreateFile(file *model.File) error {
	err := d.CreateFile(ctx, file)
	//err = d.SyncFileToElasticsearch(ctx, file, "disk_files")
	return err
}

func DeleteFile(ids []int, userID int) error {
	if len(ids) > 0 {
		err := d.DeleteFilesByIDs(ctx, ids, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsFileNameExisted(userID int, name string, parentID int) (bool, error) {
	_, err := d.GetFileByNameAndParentID(ctx, userID, name, parentID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return true, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

func GetFileList(userID, parentID, pageNum, pageSize int) ([]model.File, *int64, error) {
	files, totalSize, err := d.GetFileListByParentID(ctx, userID, parentID, pageNum, pageSize)
	return files, totalSize, err
}

func GetFileListByParentID(parentID, pageNum, pageSize int) ([]model.File, *int64, error) {
	files, totalSize, err := d.GetFileList(ctx, parentID, pageNum, pageSize)
	return files, totalSize, err
}

func GetFileListByType(userID, fileType, pageNum, pageSize int) ([]model.File, *int64, error) {
	files, totalSize, err := d.GetFileListByType(ctx, userID, fileType, pageNum, pageSize)
	return files, totalSize, err
}

func GetRecycleFileList(userID int) ([]model.File, error) {
	files, err := d.GetRecycleFileByUserID(ctx, userID)
	return files, err
}

func GetCollectFileList(userID, pageNum, pageSize int) ([]model.File, *int64, error) {
	files, totalSize, err := d.GetCollectFileByUserID(ctx, userID, pageNum, pageSize)
	return files, totalSize, err
}

func UpdateFile(file *model.File) error {
	err := d.UpdateFile(ctx, file)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFilesCollect(files, dirs []*model.File) error {
	if len(files) > 0 {
		err := d.UpdateFilesCollect(ctx, files)
		if err != nil {
			return err
		}
	}

	if len(dirs) > 0 {
		for _, dir := range dirs {
			err := d.UpdateDirectoryAndSubFilesCollectStatus(ctx, dir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateFilesMove(files []*model.File) error {
	if len(files) > 0 {
		err := d.UpdateFilesMove(ctx, files)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetFile(id, userID int) (*model.File, error) {
	file, err := d.GetFileByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileByID(id int) (*model.File, error) {
	file, err := d.GetFile(ctx, id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileExt(fileName string) string {
	idx := strings.LastIndex(fileName, ".")
	if idx == -1 {
		return ""
	}
	return fileName[idx:]
}

func GetFileSuffix(fileName string) string {
	idx := strings.LastIndex(fileName, ".")
	if idx == -1 {
		return ""
	}
	return fileName[idx+1:]
}

func JudgeFileType(fileName string) (string, int) {
	ext := strings.ToLower(GetFileSuffix(fileName))
	if ext == "" {
		return "未知文件", model.OtherType
	}

	extToType := map[string]int{
		"mp4":  model.VideoType,
		"mkv":  model.VideoType,
		"avi":  model.VideoType,
		"mov":  model.VideoType,
		"flv":  model.VideoType,
		"mp3":  model.AudioType,
		"wav":  model.AudioType,
		"flac": model.AudioType,
		"ogg":  model.AudioType,
		"txt":  model.DocumentType,
		"pdf":  model.DocumentType,
		"doc":  model.DocumentType,
		"docx": model.DocumentType,
		"xls":  model.DocumentType,
		"xlsx": model.DocumentType,
		"ppt":  model.DocumentType,
		"pptx": model.DocumentType,
		"jpg":  model.ImageType,
		"jpeg": model.ImageType,
		"png":  model.ImageType,
		"gif":  model.ImageType,
		"bmp":  model.ImageType,
		"tiff": model.ImageType,
		"webp": model.ImageType,
	}

	if fileType, found := extToType[ext]; found {
		return ext + "文件", fileType
	}

	return ext + "文件", model.OtherType
}

func FormatFileSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	s := float64(size)
	i := 0
	for s >= 1024 && i < len(units)-1 {
		s /= 1024
		i++
	}
	return fmt.Sprintf("%.2f%s", s, units[i])
}

func GetDirectoryStats(dirID int64) (*model.DirectoryStats, error) {
	stats, err := d.GetDirectoryStats(ctx, dirID)
	return stats, err
}

func RecoverFile(ids []int, userID int) error {
	if len(ids) > 0 {
		err := d.RecoverFile(ctx, ids, userID)
		return err
	}
	return nil
}

func CompleteDelFile(ids []int, userID int) error {
	if len(ids) > 0 {
		err := d.CompleteDelFile(ctx, ids, userID)
		return err
	}
	return nil
}

func StartScheduledCleanup() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() { // 每天 0 点执行
		if err := d.CleanupExpiredFiles(ctx); err != nil {
			zap.L().Warn("无法删除过期文件", zap.Error(err))
		}
	})
	if err != nil {
		zap.L().Warn("无法启动删除过期文件的定时任务", zap.Error(err))
	}
	c.Start()
}

func SetFileIDs(fileIDs []int) (string, error) {
	data, err := json.Marshal(fileIDs)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func CheckCodeValidity(code string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return re.MatchString(code)
}

func GenerateCode() (string, error) {
	randomBytes := make([]byte, 3)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	hexStr := hex.EncodeToString(randomBytes)
	code := strings.ToUpper(hexStr[:4])
	validChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte
	for _, char := range code {
		index := int(char) % len(validChars)
		if index < 0 {
			index = -index
		}
		fmt.Println(index)
		result = append(result, validChars[index])
	}
	return string(result), nil
}

func SetLinkExpiresAt(validity int) time.Time {
	var expiresAt time.Time
	switch validity {
	case 0:
		expiresAt = time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC)
	case 1:
		expiresAt = time.Now().Add(24 * time.Hour)
	case 7:
		expiresAt = time.Now().Add(7 * 24 * time.Hour)
	case 30:
		expiresAt = time.Now().Add(30 * 24 * time.Hour)
	case 365:
		expiresAt = time.Now().Add(365 * 24 * time.Hour)
	default:
		expiresAt = time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC)
	}
	return expiresAt
}

func CreateLink(link *model.ShareLink) error {
	err := d.CreateLink(ctx, link)
	return err
}

func GenerateShareLink() (string, error) {
	randomStr, err := generateRandomString(20)
	if err != nil {
		return "", err
	}
	link := fmt.Sprintf("%d%s", time.Now().Year(), randomStr)
	return link, nil
}

func generateRandomString(length int) (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		sb.WriteByte(chars[n.Int64()])
	}
	return sb.String(), nil
}

func CheckFilesExistenceByIDs(ids []int, userID int) ([]*model.File, bool, error) {
	files, existsMap, err := d.CheckFilesExistenceByIDs(ctx, ids, userID)
	if err != nil {
		return nil, false, err
	}

	for _, id := range ids {
		if !existsMap[id] {
			return nil, false, nil
		}
	}

	return files, true, nil
}

func GetLink(url string) (*model.ShareLink, error) {
	file, err := d.GetLinkByUrl(ctx, url)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileIDs(ids string) ([]int, error) {
	var fileIDs []int
	err := json.Unmarshal([]byte(ids), &fileIDs)
	return fileIDs, err
}

func GetFilesByIDs(ids []int) ([]model.File, error) {
	var files []model.File
	files, err := d.GetFilesByIDs(ctx, ids)
	return files, err
}

func StoreDirectoryContent(file *model.File, userID int, parentID int, path string) error {
	newDir := model.File{
		UserID:    userID,
		Name:      file.Name,
		Path:      path,
		Size:      0,
		FileType:  file.FileType,
		IsDir:     true,
		IsCollect: false,
		ParentID:  parentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := CreateFile(&newDir)
	if err != nil {
		return err
	}

	children, err := d.GetFilesByParentID(ctx, int(file.ID))
	if err != nil {
		return err
	}
	for _, child := range children {
		if child.IsDir {
			err := StoreDirectoryContent(&child, userID, int(file.ID), path)
			if err != nil {
				return err
			}
		} else {
			newFile := model.File{
				UserID:    userID,
				Name:      child.Name,
				Path:      path + "/" + newDir.Name,
				Size:      child.Size,
				FileType:  child.FileType,
				FileUrl:   child.FileUrl,
				IsDir:     child.IsDir,
				IsCollect: false,
				ParentID:  int(newDir.ID),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			err := CreateFile(&newFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
