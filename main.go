package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"micro/protos"
	"micro/services"
)

func main() {
	//micro v3新版本服务
	// create a new service
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	myService := micro.NewService(
		micro.Name("go.micro.api.blog"),
		micro.Address(":8001"),
		micro.Registry(etcdReg),
	)
	myService.Init()
	//注册grpc服务进来,这是服务端，需要grpc客户端调用
	protos.RegisterBlogServiceHandler(myService.Server(),services.NewBlogService("go.micro.api.blog",myService.Client()))
	// initialise flags

	// start the service
	myService.Run()
}
