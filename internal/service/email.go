package service

import (
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gopkg.in/gomail.v2"
	"math/rand"
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
