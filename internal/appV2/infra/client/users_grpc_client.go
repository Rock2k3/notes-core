package client

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	notesgrpcapi "github.com/Rock2k3/notes-grpc-api/generated-sources"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type usersGrpcClient struct {
	Conn   *grpc.ClientConn
	Client notesgrpcapi.UserServiceClient
}

func NewUsersGrpcClient() *usersGrpcClient {
	conf := config.GetAppConfig()

	conn, err := grpc.Dial(conf.Grpc().GrpcUsersServiceAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		//log.Printf("did not connect: %v", err)
	}
	client := notesgrpcapi.NewUserServiceClient(conn)

	return &usersGrpcClient{
		Conn:   conn,
		Client: client,
	}
}
