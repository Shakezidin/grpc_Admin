package user

import (
	"log"

	"github.com/shakezidin/config"
	pb "github.com/shakezidin/pkg/user/pb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCUSERPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCUSERPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Booking Client at port: %v", cfg.GRPCUSERPORT)
	return pb.NewUserServiceClient(grpc), nil
}
