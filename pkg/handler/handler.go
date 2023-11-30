package handler

import (
	"context"
	"log"

	"github.com/shakezidin/config"
	admin "github.com/shakezidin/pkg/pb/pb"
	"github.com/shakezidin/pkg/service/interfaces"
)

type AdminHandlers struct {
	cnfg         *config.Config
	AdminService interfaces.AdminServiceInter
	admin.AdminServiceServer
}

func (a *AdminHandlers) AdminLogin(ctx context.Context, p *admin.LoginRequest) (*admin.LoginResponce, error) {
	result, err := a.AdminService.AdminLogin(p, *a.cnfg)
	if err != nil {
		log.Print("Error while fetching all users")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) CreateUser(ctx context.Context, p *admin.User) (*admin.UserResponse, error) {
	result, err := a.AdminService.CreateService(p, *a.cnfg)
	if err != nil {
		log.Print("user creation error")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) DeleteUser(ctx context.Context, p *admin.DeleteUserRequest) (*admin.UserResponse, error) {
	result, err := a.AdminService.DeleteService(p, *a.cnfg)
	if err != nil {
		log.Print("Delete user error")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) SearchUser(ctx context.Context, p *admin.UserRequest) (*admin.SearchResponse, error) {
	result, err := a.AdminService.SearchUserService(p, *a.cnfg)
	if err != nil {
		log.Print("Search user error")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandlers) EditUser(ctx context.Context, p *admin.User) (*admin.UserResponse, error) {
	result, err := a.AdminService.EditUserService(p, *a.cnfg)
	if err != nil {
		log.Print("Search user error")
		return nil, err
	}
	return result, nil
}

func AdminHandler(repo interfaces.AdminServiceInter, cnfg *config.Config) *AdminHandlers {
	return &AdminHandlers{
		cnfg:         cnfg,
		AdminService: repo,
	}
}
