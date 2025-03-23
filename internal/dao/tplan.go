package dao

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
)

func (d *Dao) CreateTPlan(ctx context.Context, tPlan *model.TPlan) error {
	result := d.orm.WithContext(ctx).Model(&model.TPlan{}).Create(tPlan)
	return result.Error
}

func (d *Dao) GetTPlanByID(ctx context.Context, id int) (*model.TPlan, error) {
	var tPlan *model.TPlan
	result := d.orm.WithContext(ctx).Model(&model.TPlan{}).Where("id = ?", id).First(&tPlan)
	return tPlan, result.Error
}

func (d *Dao) UpdateTPlan(ctx context.Context, tPlan *model.TPlan) error {
	result := d.orm.WithContext(ctx).Save(&tPlan)
	return result.Error
}

func (d *Dao) GetTPlanList(ctx context.Context, userID int) ([]model.TPlan, error) {
	var tPlans []model.TPlan
	err := d.orm.WithContext(ctx).Model(&model.TPlan{}).
		Where("user_id = ?", userID).
		Order("updated_at DESC").Find(&tPlans).Error
	if err != nil {
		return nil, err
	}
	return tPlans, nil
}

func (d *Dao) DeleteTPlanByID(ctx context.Context, id, userID int) error {
	result := d.orm.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&model.TPlan{})
	return result.Error
}

func (d *Dao) UpdateTPlanUrl(ctx context.Context, id int, url string) error {
	result := d.orm.WithContext(ctx).Model(&model.TPlan{}).Where("id = ?", id).Update("t_plan_url", url)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
