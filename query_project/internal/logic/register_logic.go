package logic

import (
	"context"

	"github.com/pkg/errors"

	"QueryProject/query_project/internal/svc"
	"QueryProject/query_project/internal/types"
	"QueryProject/query_project/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	_, err = l.svcCtx.UserAdminModel.FindOneByUsername(l.ctx, req.RegisterUserName)
	if err == nil {
		logx.Error("用户名已存在")
		return nil, errors.Errorf("用户已存在")
	}
	if err != model.ErrNotFound {
		logx.Error("find user error:", err)
		return nil, err
	}

	_, err = l.svcCtx.UserAdminModel.Insert(l.ctx, &model.UserAdmin{
		Username: req.RegisterUserName,
		Passwd:   ToHmac(req.RegisterPassWord, l.svcCtx.Config.Auth.Salt),
	})
	if err != nil {
		logx.Error("insert user error:", err)
		return nil, err
	}

	return &types.RegisterResponse{
		RegisterUserName: req.RegisterUserName,
	}, nil
}
