package logic

import (
	"context"
	"go_zero-crud/rpc/models"

	"go_zero-crud/rpc/service/company_service/internal/svc"
	"go_zero-crud/rpc/service/company_service/pb/company_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCompanyLogic {
	return &CreateCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCompanyLogic) CreateCompany(in *company_service.CompanyRequest) (*company_service.EmptyResponse, error) {
	_, err := l.svcCtx.CompanyRepository.Create(&models.Company{
		Name:     in.Name,
		Password: in.Password,
		Email:    in.Email,
	})

	if err != nil {
		return nil, err
	}

	return &company_service.EmptyResponse{
		Ok: true,
	}, nil
}
