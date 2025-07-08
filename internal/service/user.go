package service

import (
	"context"
	"encoding/json"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"math/rand"
)

func CreateUser(user model.User) error {
	aesEncryptPassword(&user)
	avatar := getAvatar()
	user.Avatar = avatar
	err := d.CreateUser(ctx, &user)
	return err
}

func UpdateUser(user model.User) error {
	aesEncryptPassword(&user)
	err := d.UpdateUser(ctx, &user)
	return err
}

func GetUserByUsername(username string) (*model.User, error) {
	user, err := d.GetUserByUsername(ctx, username)
	if user.Password != "" {
		aesDecryptPassword(user)
	}
	return user, err
}

func GetUserByEmail(email string) (*model.User, error) {
	user, err := d.GetUserByEmail(ctx, email)
	if user.Password != "" {
		aesDecryptPassword(user)
	}
	return user, err
}

func GetUserByUserID(userID uint) (*model.User, error) {
	user, err := d.GetUserByUserID(ctx, userID)
	if user.Password != "" {
		aesDecryptPassword(user)
	}
	return user, err
}

func GetUserID(ctx context.Context) (int64, error) {
	userIDVal := ctx.Value("user_id")
	if userIDVal == nil {
		return 0, utils.ErrServer
	}
	jsonNumber, ok := userIDVal.(json.Number)
	if !ok {
		return 0, utils.ErrServer
	}
	userID, err := jsonNumber.Int64()
	if err != nil {
		return 0, utils.ErrServer
	}
	return userID, nil
}

func getAvatar() string {
	avatars := []string{
		"http://47.96.78.173:9000/aihelper/image/2025/b0a913fc-df81-11ef-9738-04bf1b6f946c.webp",
		"http://47.96.78.173:9000/aihelper/image/2025/ca2db7bc-df82-11ef-9c28-04bf1b6f946c.webp",
		"http://47.96.78.173:9000/aihelper/image/2025/e1647caf-df82-11ef-9c28-04bf1b6f946c.webp",
		"http://47.96.78.173:9000/aihelper/image/2025/eb7228b1-df82-11ef-9c28-04bf1b6f946c.webp",
	}
	randomIndex := rand.Intn(len(avatars))
	return avatars[randomIndex]
}

func aesDecryptPassword(user *model.User) {
	user.Password = utils.AesDecrypt(user.Password, *conf)
}

func aesEncryptPassword(user *model.User) {
	user.Password = utils.AesEncrypt(user.Password, *conf)
}
