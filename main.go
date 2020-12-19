package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"micro/Services"
	"micro/ServicesImpl"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.104:8500"),
	)

	myService := micro.NewService(
		micro.Name("micro.micro-api"),
		micro.Address(":8001"),
		micro.Registry(consulReg),
		)

	Services.RegisterTestServiceHandler(myService.Server(),new(ServicesImpl.TestService))
	myService.Run()
}
