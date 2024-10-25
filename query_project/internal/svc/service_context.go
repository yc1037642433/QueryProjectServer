package svc

import (
	"QueryProject/query_project/internal/config"
	"QueryProject/query_project/internal/middleware"
	"QueryProject/query_project/model"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config      config.Config
	Authority   rest.Middleware
	RedisClient *redis.Redis

	UserAdminModel model.UserAdminModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConf := c.DataBaseConf
	sql_source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", dbConf.UserName, dbConf.PassWord, dbConf.Host, dbConf.Port, dbConf.DBName)
	// sql_source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", "root", "yc199365", "0.0.0.0", 3306, "QueryProject")
	userAdminConn := sqlx.NewMysql(sql_source)

	RedisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		println("redis connect error")
	}

	return &ServiceContext{
		Config:         c,
		RedisClient:    RedisClient,
		Authority:      middleware.NewAuthorityMiddleware(RedisClient).Handle,
		UserAdminModel: model.NewUserAdminModel(userAdminConn),
	}
}
