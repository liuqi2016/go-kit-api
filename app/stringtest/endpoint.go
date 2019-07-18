package stringtest

//端点将每个服务方法提供为rpc接口

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndpoint     endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		UppercaseEndpoint: MakeUppercaseEndpoint(s),
	}
}

func MakeUppercaseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}
