package logic

import (
	"context"
	"go_zero-crud/rpc/models"

	"go_zero-crud/rpc/service/user_service/internal/svc"
	"go_zero-crud/rpc/service/user_service/pb/user_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_service.UserRequest) (*user_service.EmptyResponse, error) {
	if _, err := l.svcCtx.UserRepository.Create(&models.User{
		Name:     in.Name,
		Password: in.Password,
		Email:    in.Email,
	}); err != nil {
		return nil, err
	}
	return &user_service.EmptyResponse{
		Ok: true,
	}, nil
}
