package handler

import (
	"context"
	"log"

	admin "github.com/shakezidin/pkg/pb/pb"
	"github.com/shakezidin/pkg/service/interfaces"
)

type AdminHandlers struct {
	AdminService interfaces.AdminServiceInter
	admin.AdminServiceServer
}

func (a *AdminHandlers) AdminLogin(ctx context.Context, p *admin.LoginRequest) (*admin.LoginResponce, error) {
	result, err := a.AdminService.AdminLogin(p)
	if err != nil {
		log.Print("Error while fetching all users")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) CreateUser(ctx context.Context, p *admin.User) (*admin.UserResponse, error) {
	result, err := a.AdminService.CreateService(p)
	if err != nil {
		log.Print("user creation error")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) DeleteUser(ctx context.Context, p *admin.DeleteUserRequest) (*admin.UserResponse, error) {
	result, err := a.AdminService.DeleteService(p)
	if err != nil {
		log.Print("Delete user error")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) SearchUser(ctx context.Context, p *admin.UserRequest) (*admin.SearchResponse, error) {
	result, err := a.AdminService.SearchUserService(p)
	if err != nil {
		log.Print("Search user error")
		return nil, err
	}
	return result, nil
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
