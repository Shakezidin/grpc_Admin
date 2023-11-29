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

func (a *AdminServices) AdminLogin(admn *adminpb.LoginRequest, client userpb.UserServiceClient) (*adminpb.LoginResponce, error) {
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

func AdminRepository(repo Repointer.AdminRepoInter) interfaces.AdminServiceInter {
	return &AdminServices{
		adminRepo: repo,
	}
}
