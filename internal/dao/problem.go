package dao

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
)

func (d *Dao) CreateProblem(ctx context.Context, problem *model.Problem) error {
	result := d.orm.WithContext(ctx).Model(&model.Problem{}).Create(problem)
	return result.Error
}

func (d *Dao) CreateProblemsBatch(ctx context.Context, problems []*model.Problem) error {
	if err := d.orm.WithContext(ctx).Create(&problems).Error; err != nil {
		return fmt.Errorf("批量插入题目记录失败: %w", err)
	}
	return nil
}
