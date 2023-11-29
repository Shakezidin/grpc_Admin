package repository

import (
	"log"

	"github.com/shakezidin/pkg/DTO"
	"github.com/shakezidin/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AdminServices struct {
	db *gorm.DB
}

func (a *AdminServices) FetchAdmin(username string) (*DTO.Admin, error) {
	var admin DTO.Admin
	result := a.db.Where("user_name = ?", username).First(&admin)
	if result.Error != nil {
		log.Print("Error while fetching admin")
		return nil, result.Error
	}
	return &admin, nil
}

func AdminService(db *gorm.DB) interfaces.AdminRepoInter {
	return &AdminServices{
		db: db,
	}
}
