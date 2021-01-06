package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"micro/sidecar"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.Handle("POST","/test", func(context *gin.Context){
			context.JSON(200,gin.H{
				"data":"test",
			})
		})
	}
	server := &http.Server{
		Addr:              ":8088",
		Handler:           ginRouter,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	service := sidecar.NewService("api.lxtkj.cn")
	service.AddNode("test-"+uuid.New().String(),8088,"127.0.0.1:8088")
	handler := make(chan error)

	go func() {
		handler<-server.ListenAndServe()
	}()

	go func() {
		notify := make(chan os.Signal)
		signal.Notify(notify,syscall.SIGTERM,syscall.SIGINT, syscall.SIGKILL)
		handler<-fmt.Errorf("%s",<-notify)
	}()

	go func() {
		//注册服务
		err := sidecar.RegService(service)
		if err != nil {
			handler<-err
		}
	}()
	getHandler := <-handler
	fmt.Println(getHandler.Error())
	//反注册服务
	err := sidecar.UnRegService(service)
	if err != nil {
		log.Println(err)
	}
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
