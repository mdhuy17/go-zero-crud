package svc

import (
	autoMigration "go_zero-crud/rpc/auto_migration"
	infra "go_zero-crud/rpc/infra"
	"go_zero-crud/rpc/repository"
	"go_zero-crud/rpc/repository/interface_repository"
	"go_zero-crud/rpc/service/company_service/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	CompanyRepository interface_repository.CompanyDaoInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConnection := infra.ConnectDb(c.MySQLConnectionString)
	companyRepository := repository.NewCompanyRepository(dbConnection)
	if err := autoMigration.Migration(dbConnection); err != nil {
		return nil
	}
	return &ServiceContext{
		Config:            c,
		CompanyRepository: companyRepository,
	}
}
