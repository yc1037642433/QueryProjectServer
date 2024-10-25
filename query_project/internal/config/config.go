package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
		Salt          string
	}
	DataBaseConf struct {
		Host     string
		Port     int
		DBName   string
		UserName string
		PassWord string
	}
	RedisConf redis.RedisConf
}
