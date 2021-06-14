package ServicesImpl

import (
	"context"
	"micro/Services"
	"strconv"
	"time"
)

type TestService struct {

}

func(this *TestService) Call(ctx context.Context, req *Services.TestRequest, rsp *Services.TestResponse)  error {
	time.Sleep(time.Second*3)
	rsp.Data="test"+strconv.Itoa(int(req.Id))
	return nil
}