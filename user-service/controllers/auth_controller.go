package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nghiack7/micro-expense-manager/user-service/pkg/config"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/models"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/pb"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/services"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/utils"
	"gorm.io/gorm"
)

type AuthController struct {
	userService services.UserService
	pb.UnimplementedAuthServiceServer
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func (a *AuthController) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := a.userService.FindUserByEmail(req.Username)
	if err != nil {
		return nil, err
	}
	err = utils.ValidatePassword(req.Password, user.HashPassword)
	if err != nil {
		return nil, fmt.Errorf("your password is incorrect")
	}
	// Generate Tokens
	access_token, err := utils.CreateToken(config.ConfigApp.AccessTokenExpiresIn, user.ID, config.ConfigApp.AccessTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	refresh_token, err := utils.CreateToken(config.ConfigApp.RefreshTokenExpiresIn, user.ID, config.ConfigApp.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	resp := &pb.SignInResponse{
		Success:      true,
		Message:      "Login successfully",
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	return resp, nil

}

func (a *AuthController) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	newUser := &models.User{
		Username:    req.Username,
		Email:       req.Email,
		NumberPhone: req.PhoneNumber,
		Credit:      0,
	}
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("password do not match")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	newUser.HashPassword = hashedPassword
	_, err = a.userService.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	resp := &pb.SignUpResponse{
		Success: true,
		Message: fmt.Sprintf("Sign Up Success with username %s", req.Username),
	}
	return resp, nil
}

func (a *AuthController) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	sub, err := utils.ValidateToken(req.RefreshToken, config.ConfigApp.RefreshTokenPublicKey)
	if err != nil {
		return nil, err
	}
	user, err := a.userService.FindUserById(sub.(uint))
	if err != nil {
		return nil, err
	}
	access_token, err := utils.CreateToken(config.ConfigApp.AccessTokenExpiresIn, user.ID, config.ConfigApp.AccessTokenPrivateKey)
	if err != nil {
		return nil, err
	}
	resp := &pb.RefreshTokenResponse{Success: true, AccessToken: access_token}
	return resp, nil
}

func (a *AuthController) GoogleLoginOAuth(ctx context.Context, req *pb.GoogleLoginOAuthRequest) (*pb.GoogleLoginOAuthResponse, error) {

	// Use the code to get the id and access tokens
	tokenRes, err := utils.GetGoogleOauthToken(req.Code)

	if err != nil {
		return nil, err
	}

	user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		return nil, err
	}

	createdAt := time.Now()
	resBody := &models.User{
		Email:    user.Email,
		Username: user.Name,
	}
	resBody.CreatedAt = createdAt
	resBody.UpdatedAt = createdAt
	_, err = a.userService.FindUserByEmail(user.Email)
	var updatedUser *models.User
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		updatedUser, err = a.userService.UpdateUser(resBody)
		if err != nil {
			return nil, err
		}
	} else {
		updatedUser, err = a.userService.CreateUser(resBody)
		if err != nil {
			return nil, err
		}
	}

	// Generate Tokens
	access_token, err := utils.CreateToken(config.ConfigApp.AccessTokenExpiresIn, updatedUser.ID, config.ConfigApp.AccessTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	refresh_token, err := utils.CreateToken(config.ConfigApp.RefreshTokenExpiresIn, updatedUser.ID, config.ConfigApp.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}
	resp := &pb.GoogleLoginOAuthResponse{
		Success:      true,
		Message:      "Login successfully",
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	return resp, nil
}
