package logic

import (
	"context"

	"go_zero-crud/rpc/service/user_service/internal/svc"
	"go_zero-crud/rpc/service/user_service/pb/user_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user_service.UserRequest) (*user_service.UserResponse, error) {
	user, err := l.svcCtx.UserRepository.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	return &user_service.UserResponse{
		Id:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
