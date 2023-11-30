package di

import (
	"log"

	"github.com/shakezidin/config"
	"github.com/shakezidin/pkg/db"
	"github.com/shakezidin/pkg/handler"
	"github.com/shakezidin/pkg/repository"
	"github.com/shakezidin/pkg/server"
	"github.com/shakezidin/pkg/service"
	"github.com/shakezidin/pkg/user"
)

func Init() {
	config := config.LoadConfig()
	db := db.Database(config)
	client, err := user.ClientDial(config)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
	adminrepo := repository.AdminRepository(db)
	adminService := service.AdminService(adminrepo, client)
	adminHandler := handler.AdminHandler(adminService, config)
	err = server.NewGrpcServer(config, adminHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}
