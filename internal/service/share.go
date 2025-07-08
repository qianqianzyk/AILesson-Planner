package service

import (
	"bytes"
	convertapi "github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func CreateShareResource(post *model.ShareResource) error {
	err := d.CreateShareResource(ctx, post)
	return err
}

func DeleteShareResource(resourceID, userID int) error {
	coverImg, err := d.DeleteShareResource(ctx, resourceID, userID)
	if coverImg != "" {
		DeleteObjectByUrlAsync(coverImg)
	}
	return err
}

func GetShareResourceList(resourceType, contentType int) ([]model.ShareResource, error) {
	posts, err := d.GetShareResourceList(ctx, resourceType, contentType)
	return posts, err
}

func SearchShareResource(resourceType, contentType int, keyword string) ([]model.ShareResource, error) {
	posts, err := d.SearchShareResource(ctx, resourceType, contentType, keyword)
	return posts, err
}

func GetSubjectNameByName(name string) int {
	var SubjectMap = map[string]int{
		"中国近代史":   1,
		"程序设计基础C": 2,
		"线性代数B":   3,
		"大学英语":    4,
		"国家安全教育":  5,
		"高等数学":    6,
		"大学物理":    7,
	}

	if id, exists := SubjectMap[name]; exists {
		return id
	}
	return 0
}

func GetSubjectNameByID(id int) string {
	var SubjectMap = map[string]int{
		"中国近代史":   1,
		"程序设计基础C": 2,
		"线性代数B":   3,
		"大学英语":    4,
		"国家安全教育":  5,
		"高等数学":    6,
		"大学物理":    7,
	}

	for name, value := range SubjectMap {
		if value == id {
			return name
		}
	}
	return "未知科目"
}

func GetAllSubjects() []string {
	var SubjectMap = map[string]int{
		"中国近代史":   1,
		"程序设计基础C": 2,
		"线性代数B":   3,
		"大学英语":    4,
		"国家安全教育":  5,
		"高等数学":    6,
		"大学物理":    7,
	}

	subjects := make([]string, 0, len(SubjectMap))
	for name := range SubjectMap {
		subjects = append(subjects, name)
	}
	sort.Strings(subjects)
	return subjects
}

func CovertToPdf(fileUrl, apiKey string) (string, error) {
	// ConvertAPI 配置
	config.Default = config.NewDefault(apiKey)
	// MinIO 上的 PPT 文件 URL
	pptFormat := detectPPTFormat(fileUrl)
	// 调用 ConvertAPI 进行 PPT -> PDF 转换
	pdfRes := convertapi.ConvDef(pptFormat, "pdf",
		param.NewString("file", fileUrl),
	)
	// 创建一个临时文件来保存 PDF
	tempFile, err := os.CreateTemp("", "converted_*.pdf")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()
	// 使用 ToFile 方法将 PDF 保存到临时文件
	err = pdfRes.ToFile(tempFile)
	if err != nil {
		return "", err
	}
	// 读取保存的 PDF 文件内容
	pdfData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return "", err
	}
	// 构造 MinIO 上传参数
	fileSize := int64(len(pdfData))
	objectKey := GenerateObjectKey("share", ".pdf")
	// 上传 PDF 文件到 MinIO
	objectUrl, err := PutObject(objectKey, bytes.NewReader(pdfData), fileSize, "application/pdf")
	if err != nil {
		return "", err
	}
	return objectUrl, nil
}

func detectPPTFormat(fileUrl string) string {
	ext := strings.ToLower(filepath.Ext(fileUrl))
	if ext == ".pptx" {
		return "pptx"
	} else if ext == ".ppt" {
		return "ppt"
	}
	log.Fatalf("不支持的文件格式: %s", ext)
	return ""
}
