package logic

import (
	"context"
	"fmt"
	"time"

	"QueryProject/query_project/common/jwt"
	"QueryProject/query_project/internal/svc"
	"QueryProject/query_project/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user_admin, err := l.svcCtx.QueryuserModel.FindOneByUsername(l.ctx, req.LoginUserName)
	if err != nil {
		logx.Error(err)
		return nil, errors.Errorf("用户不存在")
	}
	// TODO: salt的更新机制
	hashedPassword := ToHmac(req.LoginPassWord, l.svcCtx.Config.Auth.Salt)
	if user_admin.Passwd != hashedPassword {
		logx.Error("用户名密码错误")
		return nil, errors.Errorf("用户名密码错误")
	}

	// ---------------
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	jwtToken, err := jwt.GenerateToken(
		l.svcCtx.Config.Auth.AccessSecret,
		int64(accessExpire),
		now,
		user_admin.Username,
		0,
	)
	if err != nil {
		logx.Error("生成token失败, err:", err)
		return nil, errors.Errorf("生成token失败")
	}
	// refreshToken := "refresh_token"
	redisKey := fmt.Sprintf("Token_%s", jwtToken)
	redisValue := redisKey
	err = l.svcCtx.RedisClient.Setex(redisKey, redisValue, int(accessExpire))
	if err != nil {
		logx.Error("redis set error:", err.Error())
		return nil, errors.Errorf("后端错误, token缓存失败")
	}

	// ---------------
	return &types.LoginResponse{
		LoginUserName: req.LoginUserName,
		JwtToken: types.JwtToken{
			AccessToken:  jwtToken,
			AccessExpire: now + int64(accessExpire),
			// RefreshToken: refreshToken,
		},
	}, nil
}
