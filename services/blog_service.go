package services

import (
	"context"
	"github.com/asim/go-micro/v3/client"
	"micro/protos"
	"strconv"
)

type BlogService struct {
	c    client.Client
	name string
}
func(this *BlogService) Detail(ctx context.Context, req *protos.BlogRequest, rsp *protos.BlogResponse) error{
	rsp.Data="test"+strconv.Itoa(int(req.Id))
	return nil
}

func NewBlogService(name string,c client.Client) *BlogService {
	return &BlogService{c:c,name:name}
}