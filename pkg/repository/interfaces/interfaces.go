package interfaces

import "github.com/shakezidin/pkg/DTO"

type AdminRepoInter interface {
	FetchAdmin(username string) (*DTO.Admin,error)
}
