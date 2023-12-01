package handler

import (
	"context"
	"log"

	adminpb "github.com/shakezidin/pkg/pb/pb"
	userpb "github.com/shakezidin/pkg/user/pb/pb"
)

func FetchAllSUserHandler(client userpb.AdminUserServiceClient) (*userpb.LoginResponce, error) {
	ctx := context.Background()
	responce, err := client.FetchAllSUser(ctx, &userpb.FetchUsers{})
	if err != nil {
		log.Printf("error while fetching user ", err)
		return nil, err
	}
	return responce, nil
}

func DeleteUserHandler(client userpb.AdminUserServiceClient, id uint64) (*userpb.AdminResult, error) {
	ctx := context.Background()
	responce, err := client.DeleteUser(ctx, &userpb.DeleteUserById{
		Id: id,
	})
	if err != nil {
		log.Printf("error while fetching user ", err)
		return nil, err
	}
	return responce, nil
}

func CreateUserHandler(client userpb.AdminUserServiceClient, p *adminpb.User) (*userpb.AdminResult, error) {
	ctx := context.Background()
	responce, err := client.CreateUser(ctx, &userpb.UserCreate{
		Username: p.Username,
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	})
	if err != nil {
		log.Print("Error while createing user")
		return nil, err
	}
	return responce, nil
}

func SearchUserHandler(client userpb.AdminUserServiceClient, p *adminpb.UserRequest) (*userpb.SearchResponse, error) {
	ctx := context.Background()
	responce, err := client.SearchUser(ctx, &userpb.UserRequest{
		Username: p.Username,
	})
	if err != nil {
		log.Print("Error while Searching user")
		return nil, err
	}
	return responce, nil
}

func EditUserHandler(client userpb.AdminUserServiceClient, p *adminpb.User) (*userpb.UserResponse, error) {
	ctx := context.Background()
	responce, err := client.EditUser(ctx, &userpb.Users{
		Id:       p.Id,
		Username: p.Username,
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	})
	if err != nil {
		log.Print("Error while updating user")
		return nil, err
	}
	return responce, nil
}
