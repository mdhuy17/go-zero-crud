package main

import (
	"flag"
	"fmt"

	"go_zero-crud/rpc/service/company_service/internal/config"
	"go_zero-crud/rpc/service/company_service/internal/server"
	"go_zero-crud/rpc/service/company_service/internal/svc"
	"go_zero-crud/rpc/service/company_service/pb/company_service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/companyservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		company_service.RegisterCompanyServiceServer(grpcServer, server.NewCompanyServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
