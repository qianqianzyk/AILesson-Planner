package service

import (
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gopkg.in/gomail.v2"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

const (
	SMTPServer = "smtp.qq.com"
	SMTPPort   = 465
)

func SendMsgToEmail(fromEmail, sendEmail, sendKey string, userID uint) error {
	resultStr, err := d.GetSendCodeTime(ctx, strconv.Itoa(int(userID)))
	if err != nil {
		return err
	}
	if resultStr == "" {
		err := sendVerificationCode(fromEmail, sendEmail, sendKey, userID)
		if err != nil {
			return err
		}

		err = d.StoreSendCodeTime(ctx, strconv.Itoa(int(userID)))
		if err != nil {
			return err
		}
		return nil
	} else {
		return utils.ErrTimeLimited
	}
}

func GetVerificationCode(userID uint) (string, error) {
	code, err := d.GetVerificationCode(ctx, strconv.Itoa(int(userID)))
	if err != nil {
		return "", err
	}
	return code, nil
}

func sendVerificationCode(fromEmail, sendEmail, sendKey string, userID uint) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", sendEmail)
	m.SetHeader("Subject", "[邮箱验证]")
	code := rand.Intn(900000) + 100000
	body := "验证码：" + strconv.Itoa(code) + "。有效期5分钟，请勿泄露。"
	m.SetBody("text/plain", body)
	doc := gomail.NewDialer(SMTPServer, SMTPPort, fromEmail, sendKey)
	if err := d.StoreVerificationCode(ctx, strconv.Itoa(int(userID)), strconv.Itoa(code)); err != nil {
		return err
	}
	if err := doc.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendStudentTranscripts(fromEmail, sendEmail, sendKey string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", sendEmail)
	m.SetHeader("Subject", "[XXXXXX 期末成绩单]")
	body := "你好，浅浅同学：\n  这是你2023-2024学年第1学期的期末成绩单，请注意查收！"
	m.SetBody("text/plain", body)
	// 附件Url
	attachmentURL := "https://disk.qianqianzyk.top/aihelper/transcripts/2025/xlsx/302023311111_2023-2024_1_期末成绩单.xlsx"
	// 下载附件到临时文件
	resp, err := http.Get(attachmentURL)
	if err != nil {
		return fmt.Errorf("无法下载附件: %w", err)
	}
	defer resp.Body.Close()
	// 读取文件内容
	tempFile, err := os.CreateTemp("", "302023311111_2023-2024_1_期末成绩单_temp.xlsx")
	if err != nil {
		return fmt.Errorf("创建临时文件失败: %w", err)
	}
	defer os.Remove(tempFile.Name())
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return fmt.Errorf("保存附件失败: %w", err)
	}
	// 关闭文件，防止 gomail 读取时冲突
	tempFile.Close()
	// 添加附件
	m.Attach(tempFile.Name(), gomail.Rename("302023311111_2023-2024_1_期末成绩单.xlsx"))
	// 发送邮件
	dialer := gomail.NewDialer(SMTPServer, SMTPPort, fromEmail, sendKey)
	if err := dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("发送邮件失败: %w", err)
	}
	return nil
}
