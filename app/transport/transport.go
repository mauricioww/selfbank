package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httpGokit "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(endpoints Endpoints) http.Handler {
	router := mux.NewRouter()
	router.Use(middleware)

	router.Methods("POST").Path("/goauth/create_user").Handler(httpGokit.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeResponse,
	))

	return router
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func encodeResponse(ctx context.Context, rw http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(rw).Encode(res)
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}
