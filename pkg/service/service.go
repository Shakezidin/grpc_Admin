package service

import (
	"errors"
	"log"

	"github.com/shakezidin/config"
	"github.com/shakezidin/pkg/JWT"
	adminpb "github.com/shakezidin/pkg/pb/pb"
	Repointer "github.com/shakezidin/pkg/repository/interfaces"
	"github.com/shakezidin/pkg/service/interfaces"
	user "github.com/shakezidin/pkg/user/handler"
	userpb "github.com/shakezidin/pkg/user/pb/pb"
)

type AdminServices struct {
	adminRepo Repointer.AdminRepoInter
	client    userpb.AdminUserServiceClient
}

func (a *AdminServices) AdminLogin(admn *adminpb.LoginRequest, cnfg config.Config) (*adminpb.LoginResponce, error) {
	admin, err := a.adminRepo.FetchAdmin(admn.Username)
	if err != nil {
		return nil, err
	}
	if admin.Password != admn.Password {
		log.Print("Password error")
		return nil, errors.New("password error")
	}
	result, err := user.FetchAllSUserHandler(a.client)
	if err != nil {
		log.Print("fethcing error")
		return nil, err
	}

	token, err := JWT.GenerateJWT(admn.Username, "admin")
	if err != nil {
		log.Print("Generate jwt error")
		return nil, err
	}
	var users []*adminpb.User
	for _, r := range result.Available {
		users = append(users, &adminpb.User{
			Id:       r.Id,
			Username: r.Username,
			Name:     r.Name,
			Email:    r.Email,
			Password: r.Password,
		})
	}
	rstl := &adminpb.LoginResponce{
		Status:    "success",
		Available: users,
		Token:     token,
	}
	return rstl, nil
}

func (a *AdminServices) CreateService(p *adminpb.User, cnfg config.Config) (*adminpb.UserResponse, error) {

	result, err := user.CreateUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	rslt := &adminpb.UserResponse{
		Status:   result.Status,
		Username: result.Message,
	}
	return rslt, nil
}

func (a *AdminServices) DeleteService(p *adminpb.DeleteUserRequest, cnfg config.Config) (*adminpb.UserResponse, error) {
	result, err := user.DeleteUserHandler(a.client, p.Id)
	if err != nil {
		return nil, err
	}
	rslt := &adminpb.UserResponse{
		Status:   result.Status,
		Username: result.Message,
	}
	return rslt, nil
}

func (a *AdminServices) SearchUserService(p *adminpb.UserRequest, cnfg config.Config) (*adminpb.SearchResponse, error) {
	result, err := user.SearchUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	var users []*adminpb.User
	for _, i := range result.Available {
		users = append(users, &adminpb.User{
			Id:       i.Id,
			Username: i.Username,
			Name:     i.Name,
			Email:    i.Email,
			Password: i.Password,
		})
	}
	rslt := &adminpb.SearchResponse{
		Status:    result.Status,
		Available: users,
	}
	return rslt, nil
}

func (a *AdminServices) EditUserService(p *adminpb.User, cnfg config.Config) (*adminpb.UserResponse, error) {
	result, err := user.EditUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	rslt := &adminpb.UserResponse{
		Status:   result.Status,
		Username: result.Username,
	}
	return rslt, nil
}

func AdminService(repo Repointer.AdminRepoInter, client userpb.AdminUserServiceClient) interfaces.AdminServiceInter {
	return &AdminServices{
		adminRepo: repo,
		client:    client,
	}
}
