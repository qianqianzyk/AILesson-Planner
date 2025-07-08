package dao

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
)

func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	result := d.orm.WithContext(ctx).Model(&model.User{}).Create(user)
	return result.Error
}

func (d *Dao) UpdateUser(ctx context.Context, user *model.User) error {
	result := d.orm.WithContext(ctx).Save(&user)
	return result.Error
}

func (d *Dao) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	result := d.orm.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (d *Dao) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := d.orm.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (d *Dao) GetUserByUserID(ctx context.Context, userID uint) (*model.User, error) {
	var user model.User
	result := d.orm.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).First(&user)
	return &user, result.Error
}
