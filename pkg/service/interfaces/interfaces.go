package interfaces

import (
	adminpb "github.com/shakezidin/pkg/pb/pb"
)

type AdminServiceInter interface {
	AdminLogin(admn *adminpb.LoginRequest) (*adminpb.LoginResponce, error)
	CreateService(p *adminpb.User) (*adminpb.UserResponse, error)
	DeleteService(p *adminpb.DeleteUserRequest)(*adminpb.UserResponse,error)
	SearchUserService(p *adminpb.UserRequest) (*adminpb.SearchResponse, error)
}
