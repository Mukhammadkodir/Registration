package services

import (
	"fmt"
	"github/RegisterTask/Api/config"
	pbu "github/RegisterTask/Api/genproto/register_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	RegisterService() pbu.RegisterServiceClient
}

type serviceManager struct {
	registerService pbu.RegisterServiceClient
}


func (s *serviceManager) RegisterService() pbu.RegisterServiceClient {
	return s.registerService
}


func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.RegisterServiceHost, conf.RegisterServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		registerService: pbu.NewRegisterServiceClient(connUser),
	}

	return serviceManager, nil
}
