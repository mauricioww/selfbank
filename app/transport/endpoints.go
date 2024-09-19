package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mauricioww/goauth/app/service"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeHttpEndpoints(s service.Servicer) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		res, err := s.CreateUser(req.FirstName, req.LastName, req.Email, req.Password, req.Admin)
		return CreateUserResponse{Success: res}, err
	}
}
