package handlers

import (
	"context"
	"errors"

	authv1 "github.com/MMaxaMM/ChatProject/Auth/gen"
	"github.com/MMaxaMM/ChatProject/Auth/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthGRPC interface {
	Login(
		ctx context.Context,
		username string,
		password string,
	) (token string, err error)
	Register(
		ctx context.Context,
		username string,
		password string,
	) (userID int64, err error)
}

type serverAPI struct {
	authv1.UnimplementedAuthServer
	auth AuthGRPC
}

func Register(gRPC *grpc.Server, auth AuthGRPC) {
	authv1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {

	token, err := s.auth.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authv1.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *authv1.RegisterRequest,
) (*authv1.RegisterResponse, error) {

	userID, err := s.auth.Register(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, services.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authv1.RegisterResponse{UserId: userID}, nil
}
