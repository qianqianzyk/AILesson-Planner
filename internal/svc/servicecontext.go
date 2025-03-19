package svc

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-resty/resty/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/qianqianzyk/AILesson-Planner/internal/config"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service/ws"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MinioService struct {
	Client  *minio.Core
	Bucket  string
	Domain  string
	TempDir string
}

type ServiceContext struct {
	Config           config.Config
	MySQLClient      *gorm.DB
	RedisClient      *redis.Redis
	ESClient         *elasticsearch.Client
	WebsocketManager *ws.WebSocketManager
	MinioClient      *MinioService
	RestyClient      *resty.Client
	Neo4jDriver      *neo4j.DriverWithContext
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 MySQL 客户端
	mysqlClient, err := gorm.Open(mysql.Open(c.MySQL.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to MySQL: %v", err))
	}

	if err := autoMigrate(mysqlClient); err != nil {
		panic(fmt.Sprintf("database migrate failed: %w", err))
	}

	// 初始化 Redis 客户端
	redisConf := redis.RedisConf{
		Host: c.Redis.Host,
		Pass: c.Redis.Password,
		Type: c.Redis.Type,
		Tls:  c.Redis.Tls,
	}
	// 使用 MustNewRedis 来初始化 Redis 客户端
	redisClient := redis.MustNewRedis(redisConf)

	// 初始化 Elasticsearch 客户端
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:    []string{c.ES.Host},
		Username:     c.ES.Username,
		Password:     c.ES.Password,
		DisableRetry: c.ES.DisableRetry,
		MaxRetries:   c.ES.MaxRetries,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create Elasticsearch client: %v", err))
	}

	// 初始化 Websocket 管理
	websocketManager := ws.NewWebSocketManager()

	// 初始化 Minio 客户端
	minioClient, err := minio.NewCore(c.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.SecretKey, ""),
		Secure: c.Minio.Secure,
	})
	if err != nil {
		panic(fmt.Sprintf("minio initialization failed: %w", err))
	}
	minioService := MinioService{
		Client:  minioClient,
		Bucket:  c.Minio.Bucket,
		Domain:  c.Minio.Domain,
		TempDir: c.Minio.TempDir,
	}

	// 初始化 Neo4j
	neo4jDriver, err := neo4j.NewDriverWithContext(
		c.Neo4j.Url,
		neo4j.BasicAuth(c.Neo4j.Username, c.Neo4j.Password, ""),
		func(config *neo4j.Config) {
			config.MaxConnectionPoolSize = 10
			config.Log = neo4j.ConsoleLogger(neo4j.INFO)
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create Neo4j driver: %v", err))
	}

	return &ServiceContext{
		Config:           c,
		MySQLClient:      mysqlClient,
		RedisClient:      redisClient,
		ESClient:         esClient,
		WebsocketManager: websocketManager,
		MinioClient:      &minioService,
		RestyClient:      resty.New(),
		Neo4jDriver:      &neo4jDriver,
	}
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.ConversationSession{},
		&model.ConversationMessage{},
		&model.File{},
		&model.Attachment{},
		&model.ShareLink{},
		&model.Course{},
		&model.Student{},
		&model.Score{},
		&model.ChapterScore{},
		&model.Problem{},
		&model.WrongProblem{},
	)
}
