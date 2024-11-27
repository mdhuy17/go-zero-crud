package logic

import (
	"context"
	"go_zero-crud/rpc/service/user_service/pb/user_service"

	"go_zero-crud/internal/svc"
	"go_zero-crud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserRequest) (resp *types.UserResponse, err error) {
	user, err := l.svcCtx.UserService.GetUser(l.ctx, &user_service.UserRequest{
		Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
