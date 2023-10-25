/*
This service is responsible for user authentication and management
It provides access to the application with token or password authentication
Allow the user to use github or google for authentication
Learning english is production, which is recommended for most users as a student or business to learn and develop languages skill level
*/

package main

import (
	"log"
	"net"

	"github.com/nghiack7/micro-expense-manager/user-service/controllers"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/pb"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userService := services.NewUserServices()
	userC := controllers.NewUserController(userService)
	authC := controllers.NewAuthController(userService)
	pb.RegisterUserServiceServer(s, userC)
	pb.RegisterAuthServiceServer(s, authC)

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
