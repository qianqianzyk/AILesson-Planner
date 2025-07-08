package service

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-resty/resty/v2"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/qianqianzyk/AILesson-Planner/internal/config"
	"github.com/qianqianzyk/AILesson-Planner/internal/dao"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

var (
	ctx  = context.Background()
	conf *config.Config
	d    *dao.Dao
)

func ServiceInit(db *gorm.DB, rc *redis.Redis, es *elasticsearch.Client, mn *svc.MinioService, rs *resty.Client, ne *neo4j.DriverWithContext) {
	d = dao.NewDao(db, rc, es, mn, rs, ne)
}

func ConfigInit(c *config.Config) {
	conf = c
}
