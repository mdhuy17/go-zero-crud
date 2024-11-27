package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero-crud/internal/config"
	"go_zero-crud/rpc/service/company_service/companyservice"
	"go_zero-crud/rpc/service/user_service/userservice"
)

type ServiceContext struct {
	Config         config.Config
	UserService    userservice.UserService
	CompanyService companyservice.CompanyService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserService:    userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
		CompanyService: companyservice.NewCompanyService(zrpc.MustNewClient(c.CompanyService)),
	}
}
