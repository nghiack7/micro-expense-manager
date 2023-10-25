package controllers

import (
	"context"

	"github.com/nghiack7/micro-expense-manager/user-service/pkg/pb"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/services"
)

type AuthController struct {
	userService services.UserService
	pb.UnimplementedAuthServiceServer
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func (a *AuthController) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	panic("not implemented")
}

func (a *AuthController) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	panic("not implemented")
}
