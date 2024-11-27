package logic

import (
	"context"

	"go_zero-crud/rpc/service/company_service/internal/svc"
	"go_zero-crud/rpc/service/company_service/pb/company_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanyLogic {
	return &GetCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCompanyLogic) GetCompany(in *company_service.CompanyRequest) (*company_service.CompanyResponse, error) {
	user, err := l.svcCtx.CompanyRepository.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	return &company_service.CompanyResponse{
		Id:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
