package middleware

import (
	"QueryProject/query_project/common/result"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type AuthorityMiddleware struct {
	RedisClient *redis.Redis
}

func NewAuthorityMiddleware(RedisClient *redis.Redis) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		RedisClient: RedisClient,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// check jwt token from redis
		token := r.Header.Get("Authorization")
		// 去掉token前面的Bearer
		token = strings.TrimPrefix(token, "Bearer ")

		redisKey := "Token_" + token
		redisValue := redisKey
		res, err := m.RedisClient.Get(redisKey)
		if err != nil {
			result.AuthHttpResult(r, w, nil, err)
			return
		}
		if res != redisValue {
			result.AuthHttpResult(r, w, nil, errors.Errorf("token is invalid"))
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
