package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"micro/Services"
	"micro/ServicesImpl"
)

func main() {
	//micro v3新版本服务
	// create a new service
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	myService := micro.NewService(
		micro.Name("gw.100txy.com.blog"),
		micro.Address(":8001"),
		micro.Registry(etcdReg),
	)
	//注册grpc服务进来,这是服务端，需要grpc客户端调用
	Services.RegisterTestServiceHandler(myService.Server(),new(ServicesImpl.TestService))
	// initialise flags
	myService.Init()
	// start the service
	myService.Run()
}
