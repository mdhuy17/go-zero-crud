package logic

import (
	"context"
	"go_zero-crud/rpc/service/user_service/pb/user_service"

	"go_zero-crud/internal/svc"
	"go_zero-crud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserRequest) (resp *types.Empty, err error) {
	if _, err = l.svcCtx.UserService.CreateUser(l.ctx, &user_service.UserRequest{
		Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}); err != nil {
		return nil, err
	}
	return &types.Empty{Ok: true}, nil
}
