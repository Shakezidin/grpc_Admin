package handler

import (
	"context"
	"log"

	admin "github.com/shakezidin/pkg/pb/pb"
	user "github.com/shakezidin/pkg/user/pb/pb"
)

func FetchAllSUser(client user.UserServiceClient) (*user.LoginResponce, error) {
	ctx := context.Background()
	responce, err := client.FetchAllSUser(ctx, &user.FetchUsers{})
	if err != nil {
		log.Printf("error while fetching user ", err)
		return nil, err
	}
	return responce, nil
}

func DeleteUser(client user.UserServiceClient, username string) (*user.Result, error) {
	ctx := context.Background()
	responce, err := client.DeleteUser(ctx, &user.UserRequest{
		Username: username,
	})
	if err != nil {
		log.Printf("error while fetching user ", err)
		return nil, err
	}
	return responce, nil
}

func CreateUser(client user.UserServiceClient, p *admin.User) (*user.Result, error) {
	ctx := context.Background()
	responce, err := client.CreateUser(ctx, &user.UserCreate{
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
