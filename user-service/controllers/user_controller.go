package controllers

import (
	"context"

	"github.com/nghiack7/micro-expense-manager/user-service/pkg/models"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/pb"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/services"
)

type UserController struct {
	userService services.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetCurrentUser(req *pb.GetCurrentUserRequest, stream pb.UserService_GetCurrentUserServer) error {
	user, err := c.userService.FindUserById(uint(req.Id))
	if err != nil {
		return err
	}
	resp := &pb.GetCurrentUserResponse{
		Code: 200,
		User: &pb.User{
			Username:       user.Username,
			Email:          user.Email,
			PasswordHashed: user.HashPassword,
			NumberPhone:    user.NumberPhone,
			Credit:         uint32(user.Credit),
		},
	}
	if err := stream.Send(resp); err != nil {
		return err
	}
	return nil
}

func (c *UserController) UpdateCurrentUser(ctx context.Context, req *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	user := models.User{
		Username:     req.User.Username,
		Email:        req.User.Email,
		HashPassword: req.User.PasswordHashed,
		NumberPhone:  req.User.NumberPhone,
	}
	user.ID = uint(req.Id)
	newUser, err := c.userService.UpdateUser(&user)
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateCurrentUserResponse{
		Code: 200,
		User: &pb.User{
			Username:       newUser.Username,
			Email:          newUser.Email,
			NumberPhone:    newUser.NumberPhone,
			PasswordHashed: newUser.HashPassword,
			Credit:         uint32(newUser.Credit),
		},
	}
	return resp, nil
}
