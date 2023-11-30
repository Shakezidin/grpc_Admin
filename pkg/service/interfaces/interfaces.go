package interfaces

import (
	"github.com/shakezidin/config"
	adminpb "github.com/shakezidin/pkg/pb/pb"
)

type AdminServiceInter interface {
	AdminLogin(admn *adminpb.LoginRequest, cnfg config.Config) (*adminpb.LoginResponce, error)
	CreateService(p *adminpb.User, cnfg config.Config) (*adminpb.UserResponse, error)
	DeleteService(p *adminpb.DeleteUserRequest, cnfg config.Config) (*adminpb.UserResponse, error)
	SearchUserService(p *adminpb.UserRequest, cnfg config.Config) (*adminpb.SearchResponse, error)
	EditUserService(p *adminpb.User, cnfg config.Config) (*adminpb.UserResponse, error)
}
