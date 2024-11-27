package logic

import (
	"context"

	"go_zero-crud/internal/svc"
	"go_zero-crud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go_zero-crud/rpc/service/company_service/pb/company_service"
)

type GetCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanyLogic {
	return &GetCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompanyLogic) GetCompany(req *types.CompanyRequest) (resp *types.CompanyResponse, err error) {
	company, err := l.svcCtx.CompanyService.GetCompany(l.ctx, &company_service.CompanyRequest{
		Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.CompanyResponse{
		Id:    company.Id,
		Name:  company.Name,
		Email: company.Email,
	}, nil
}
