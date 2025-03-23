package config

import (
	"github.com/zeromicro/go-zero/rest"
)

const ISSUER = "AILesson-Planner-api"

type Config struct {
	Host     string
	Port     int
	RestConf rest.RestConf
	Log      LogConfig
	Tongyi   TongyiConfig
	ES       ElasticsearchConfig
	Minio    MinioConfig

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Aes    AesConfig
	Email  EmailConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
	Neo4j  Neo4jConfig
	AI     AIConfig
	Unidoc UnidocConfig
}

type AesConfig struct {
	Key string
}

type EmailConfig struct {
	Name string
	Key  string
}

type TongyiConfig struct {
	Endpoint string
	APIKey   string
}

type ElasticsearchConfig struct {
	Host         string
	Username     string
	Password     string
	Index        string
	DisableRetry bool
	MaxRetries   int
}

type MySQLConfig struct {
	DSN string
}

type RedisConfig struct {
	Host     string
	Password string
	Type     string
	Tls      bool
}

type MinioConfig struct {
	EndPoint  string
	AccessKey string
	SecretKey string
	Secure    bool
	Bucket    string
	Domain    string
	TempDir   string
}

type LogConfig struct {
	DisableStacktrace bool
	Level             string
	Name              string
	Writer            string
	LoggerDir         string
	LogMaxSize        int
	LogMaxAge         int
	LogCompress       bool
}

type Neo4jConfig struct {
	Url            string
	Username       string
	Password       string
	AuraInstanceID string
}

type AIConfig struct {
	ChatEndpoint  string
	TPlanEndpoint string
}

type UnidocConfig struct {
	APIKey string
}
