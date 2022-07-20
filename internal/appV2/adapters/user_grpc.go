package adapters

import (
	"context"
	"github.com/Rock2k3/notes-core/internal/appV2/domain/users"
	"github.com/Rock2k3/notes-core/internal/appV2/infra/client"
	notesgrpcapi "github.com/Rock2k3/notes-grpc-api/generated-sources"
	"github.com/google/uuid"
	"time"
)

type usersGrpcAdapter struct {
}

func NewUsersGrpcAdapter() *usersGrpcAdapter {
	return &usersGrpcAdapter{}
}

func (a *usersGrpcAdapter) GetUserByUUID(uuid uuid.UUID) (*users.MyUser, error) {
	usersGrpcClient := client.NewUsersGrpcClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer usersGrpcClient.Conn.Close()
	defer cancel()

	resp, err := usersGrpcClient.Client.GetUser(ctx, &notesgrpcapi.GetUserRequest{UserId: uuid.String()})
	if err != nil {
		return nil, err
	}

	user := &users.MyUser{
		UserId: uuid,
		Name:   resp.GetUser().Name,
	}

	return user, nil
}

func (a *usersGrpcAdapter) GetUserByName(name string) (*users.MyUser, error) {
	//usersGrpcClient := NewUsersGrpcClient()
	//defer usersGrpcClient.conn.Close()
	//defer usersGrpcClient.usersGrpcClientContext.cancel()
	//
	//resp, err := usersGrpcClient.client.AddUser(usersGrpcClient.usersGrpcClientContext.ctx, &notesgrpcapi.AddUserRequest{Name: name})
	//if err != nil {
	//	return nil, err
	//}
	//
	//userId, _ := uuid.Parse(resp.GetUser().UserId)
	//user := &users.MyUser{
	//	UserId: userId,
	//	Name:   name,
	//}
	//
	//return user, nil
	return nil, nil
}

func (a *usersGrpcAdapter) AddUser(name string) (*users.MyUser, error) {
	//usersGrpcClient := NewUsersGrpcClient()
	//defer usersGrpcClient.conn.Close()
	//defer usersGrpcClient.usersGrpcClientContext.cancel()
	//
	//resp, err := usersGrpcClient.client.AddUser(usersGrpcClient.usersGrpcClientContext.ctx, &notesgrpcapi.AddUserRequest{Name: name})
	//if err != nil {
	//	return nil, err
	//}
	//
	//userId, _ := uuid.Parse(resp.GetUser().UserId)
	//user := &users.MyUser{
	//	UserId: userId,
	//	Name:   name,
	//}
	//
	//return user, nil
	return nil, nil
}
