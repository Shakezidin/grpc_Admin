package di

import (
	"log"

	"github.com/shakezidin/config"
	"github.com/shakezidin/pkg/db"
	"github.com/shakezidin/pkg/handler"
	"github.com/shakezidin/pkg/repository"
	"github.com/shakezidin/pkg/server"
	"github.com/shakezidin/pkg/service"
)

func Init() {
	config := config.LoadConfig()
	db := db.Database(config)
	adminrepo := repository.AdminService(db)
	adminService := service.AdminRepository(adminrepo)
	adminHandler := handler.AdminHandler(adminService)
	err := server.NewGrpcServer(config, adminHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}

}
