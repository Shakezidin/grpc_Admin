package handler

import (
	"context"
	"log"

	admin "github.com/shakezidin/pkg/pb/pb"
	"github.com/shakezidin/pkg/service/interfaces"
	user "github.com/shakezidin/pkg/user/handler"
	userpb "github.com/shakezidin/pkg/user/pb/pb"
)

type AdminHandlers struct {
	AdminService interfaces.AdminServiceInter
	admin.AdminServiceServer
}

func (a *AdminHandlers) AdminLogin(ctx context.Context, p *admin.LoginRequest) (*admin.LoginResponce, error) {
	var client userpb.UserServiceClient
	result, err := a.AdminService.AdminLogin(p, client)
	if err != nil {
		log.Print("Error while fetching all users")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) CreateUser(ctx context.Context, p *admin.User) (*admin.UserResponse, error) {
	var client userpb.UserServiceClient
	result, err := user.CreateUser(client, p)
	if err != nil {
		log.Print("user creation error")
		return nil, err
	}
	rslt := &admin.UserResponse{
		Status:   result.Status,
		Username: result.Message,
	}
	return rslt, nil
}

func AdminHandler(repo interfaces.AdminServiceInter) *AdminHandlers {
	return &AdminHandlers{
		AdminService: repo,
	}
}

// AdminLogin(context.Context, *LoginRequest) (*LoginResponce, error)
// 	CreateUser(context.Context, *User) (*UserResponse, error)
// 	DeleteUser(context.Context, *DeleteUserRequest) (*UserResponse, error)
// 	SearchUser(context.Context, *UserRequest) (*SearchResponse, error)
// 	EditUser(context.Context, *User) (*UserResponse, error)
