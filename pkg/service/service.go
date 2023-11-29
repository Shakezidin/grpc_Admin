package service

import (
	"errors"
	"log"

	adminpb "github.com/shakezidin/pkg/pb/pb"
	Repointer "github.com/shakezidin/pkg/repository/interfaces"
	"github.com/shakezidin/pkg/service/interfaces"
	user "github.com/shakezidin/pkg/user/handler"
	userpb "github.com/shakezidin/pkg/user/pb/pb"
)

type AdminServices struct {
	adminRepo Repointer.AdminRepoInter
}

func (a *AdminServices) AdminLogin(admn *adminpb.LoginRequest) (*adminpb.LoginResponce, error) {
	var client userpb.UserServiceClient
	admin, err := a.adminRepo.FetchAdmin(admn.Username)
	if err != nil {
		return nil, err
	}
	if admin.Password != admn.Password {
		log.Print("Password error")
		return nil, errors.New("password error")
	}
	result, err := user.FetchAllSUser(client)
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
		Token:     result.Token,
	}
	return rstl, nil
}

func (a *AdminServices) CreateService(p *adminpb.User) (*adminpb.UserResponse, error) {
	var client userpb.UserServiceClient
	result, err := user.CreateUser(client, p)
	if err != nil {
		return nil, err
	}
	rslt := &adminpb.UserResponse{
		Status:   result.Status,
		Username: result.Message,
	}
	return rslt, nil
}

func (a *AdminServices) DeleteService(p *adminpb.DeleteUserRequest) (*adminpb.UserResponse, error) {
	var client userpb.UserServiceClient
	result, err := user.DeleteUser(client, p.Id)
	if err != nil {
		return nil, err
	}
	rslt := &adminpb.UserResponse{
		Status:   result.Status,
		Username: result.Message,
	}
	return rslt, nil
}

func (a *AdminServices) SearchUserService(p *adminpb.UserRequest) (*adminpb.SearchResponse, error) {
	var client userpb.UserServiceClient
	result, err := user.SearchUser(client, p)
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

func AdminRepository(repo Repointer.AdminRepoInter) interfaces.AdminServiceInter {
	return &AdminServices{
		adminRepo: repo,
	}
}
