package interfaces

import (
	adminpb "github.com/shakezidin/pkg/pb/pb"
	userpb "github.com/shakezidin/pkg/user/pb/pb"
)

type AdminServiceInter interface {
	AdminLogin(admn *adminpb.LoginRequest, client userpb.UserServiceClient) (*adminpb.LoginResponce, error)
}
