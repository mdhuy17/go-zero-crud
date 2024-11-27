package logic

import (
	"context"
	"go_zero-crud/rpc/service/company_service/companyservice"

	"go_zero-crud/internal/svc"
	"go_zero-crud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCompanyLogic {
	return &CreateCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCompanyLogic) CreateCompany(req *types.CompanyRequest) (resp *types.Empty, err error) {
	if _, err = l.svcCtx.CompanyService.CreateCompany(l.ctx, &companyservice.CompanyRequest{
		Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}); err != nil {
		return nil, err
	}
	return &types.Empty{Ok: true}, nil
}
