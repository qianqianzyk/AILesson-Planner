package dao

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-resty/resty/v2"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type Dao struct {
	orm *gorm.DB
	rc  *redis.Redis
	es  *elasticsearch.Client
	mn  *svc.MinioService
	rs  *resty.Client
	ne  *neo4j.DriverWithContext
}

func NewDao(orm *gorm.DB, rc *redis.Redis, es *elasticsearch.Client, mn *svc.MinioService, rs *resty.Client, ne *neo4j.DriverWithContext) *Dao {
	return &Dao{
		orm: orm,
		rc:  rc,
		es:  es,
		mn:  mn,
		rs:  rs,
		ne:  ne,
	}
}
