package adapters

import (
	"context"
	NotesGrpcApi "github.com/Rock2k3/notes-grpc-api/generated-sources"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"notes-core/internal/config"
	"time"
)

type usersGrpcClientContext struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type usersGrpcClient struct {
	conn                   *grpc.ClientConn
	client                 NotesGrpcApi.UserServiceClient
	usersGrpcClientContext usersGrpcClientContext
}

func NewUsersGrpcClient() *usersGrpcClient {
	conf, _ := config.NewAppConfig()

	conn, err := grpc.Dial(conf.GrpcUsersServiceUrl(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	client := NotesGrpcApi.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return &usersGrpcClient{
		conn:   conn,
		client: client,
		usersGrpcClientContext: usersGrpcClientContext{
			ctx:    ctx,
			cancel: cancel,
		},
	}
}
