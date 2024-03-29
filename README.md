# micro
micro是一个基于micro工具集框架开发的grpc服务项目，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点，让你快速了解go微服务项目

## 说明
micro项目使用go-micro v3版本 ，grpc的服务端，服务名称：go.micro.blog 监听端口为：8001  
grpc之间的调用相当于服务之间的调用，微服务还需要统一网关服务  
这个grpc服务项目(micor)需要注册到etcd服务中去，而调用它的服务(micro-cli：对外提供接口去中转调用micro grpc服务)是不需要注册到服务中心去的，因为它有可能是外部服务在调用，当然如果是内部服务也需要注册到注册中心去,再nginx做反向代理限流等  

## 项目中初始化文件
- **启动etct服务：** 客户端启动etcd服务
- **web.bat**：通过micro工具构建etcd服务中心web界面,默认监听端口8082， Transport [http] Listening on [::]:58130端口未指定，默认服务只有go.micro.web和go.micro.http.broker
- **gw.bat**：用于构建micro api网关服务,默认监听端口8080， [api] Server [grpc] Listening on [::]:58258端口未指定，请求：127.0.0.1:8080/blog/TestService/Call 【网关空间go.micro.api，项目启动名称go.micro.api.blog】
- **gen.bat**：用于构建protobuf中间文件
- **reg.bat**：用于集成外部api到micro体系中
- **RabbitMQ管理页面：** 用户：admin，密码：mogu2018
- **Nacos管理页面：** 用户：nacos，密码：nacos
- **Sentinel管理页面：** 用户：sentinel，密码：sentinel
- **监控页面**：用户：user，密码：password123


## 项目启动
**web.bat**：启动micro web界面服务
```go
2021-06-12 16:09:07.550288 I | [web] HTTP API Listening on [::]:8082
2021-06-12 16:09:07.551291 I | [web] Transport [http] Listening on [::]:58130
2021-06-12 16:09:07.551291 I | [web] Broker [http] Connected to [::]:58131
2021-06-12 16:09:07.571344 I | [web] Registry [etcd] Registering node: go.micro.web-6feb1024-a20a-4c4c-94f9-966a5d401d3b

```

**gw.bat**：启动网关服务
```go
2021-06-12 16:16:49.348338 I | [api] Registering API Default Handler at /
2021-06-12 16:16:49.349342 I | [api] HTTP API Listening on [::]:8080
2021-06-12 16:16:49.355356 I | [api] Server [grpc] Listening on [::]:58258
2021-06-12 16:16:49.356361 I | [api] Broker [http] Connected to [::]:58259
2021-06-12 16:16:49.375447 I | [api] Registering node: go.micro.api-cf4e6fa5-2f61-4bdb-b33e-3edd077811a5
```


**go run main.go**：启动gw.100txy.com.blog服务
```go
2021-08-18 14:34:51  file=v3/service.go:192 level=info Starting [service] gw.100txy.com.blog
2021-08-18 14:34:51  file=server/rpc_server.go:820 level=info service=server Transport [http] Listening on [::]:8001
2021-08-18 14:34:51  file=server/rpc_server.go:840 level=info service=server Broker [http] Connected to 127.0.0.1:59974
2021-08-18 14:34:52  file=server/rpc_server.go:654 level=info service=server Registry [etcd] Registering node: gw.100txy.com.blog-3bf34509-b74f-4742-aac7-ee27f8c00c1d
```