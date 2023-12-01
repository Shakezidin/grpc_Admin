package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shakezidin/config"
	"github.com/shakezidin/pkg/handler"
	adminpb "github.com/shakezidin/pkg/pb/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(config *config.Config, handler *handler.AdminHandlers) error{
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", config.GRPCADMINRPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	grp := grpc.NewServer()
	adminpb.RegisterAdminServiceServer(grp, handler)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}

	log.Printf("listening on gRPC server %v", config.GRPCADMINRPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
