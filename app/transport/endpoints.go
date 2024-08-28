package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mauricioww/goauth/app/service"
)

type HttpEndpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeHttpEndpoints(s service.Service) HttpEndpoints {
	return HttpEndpoints{
		CreateUser: makeCreateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		res, err := s.CreateUser(req.FirstName, req.LastName, req.Password, req.Admin)
		return CreateUserResponse{Success: res}, err
	}
}
