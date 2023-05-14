package domain

import (
	"context"
	"fmt"
	"google.golang.org/grpc/resolver"
	"log"

	"github.com/morris-zheng/go-slim-micro-api/internal/conf"

	"github.com/morris-zheng/go-slim-core/discovery"
	"github.com/morris-zheng/go-slim-core/logger"
	"github.com/morris-zheng/go-slim-micro-usersvc/export/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceContext struct {
	Config  *conf.Config
	Logger  logger.Logger
	UserCli user.ServiceClient
}

var svc *ServiceContext

func NewServiceContext(c *conf.Config) *ServiceContext {
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	}

	if svc == nil {
		// logger
		l, err := logger.NewLogger(logger.Level(c.Logger.Level))
		if err != nil {
			log.Fatal("logger init err: ", err)
		}

		svc = &ServiceContext{
			Config: c,
			Logger: l,
		}
	}

	initUserCli(svc, grpcOpts)

	return svc
}

func initUserCli(svc *ServiceContext, opts []grpc.DialOption) {
	addr := ""
	etcdConfig := svc.Config.Etcd

	if len(etcdConfig.Endpoints) != 0 {
		etcdResolver := discovery.NewResolver(discovery.Option{
			Endpoints: etcdConfig.Endpoints,
			Prefix:    etcdConfig.Prefix,
		}, &svc.Logger)

		addr = discovery.Scheme("usersvc")
		resolver.Register(etcdResolver)
	}

	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		svc.Logger.Fatal(context.Background(), fmt.Sprintf("init user client error: %v", err))
	}

	svc.UserCli = user.NewServiceClient(cc)
}
