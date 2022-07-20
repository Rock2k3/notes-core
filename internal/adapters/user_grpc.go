package adapters

import (
	"context"
	"github.com/Rock2k3/notes-core/internal/domain/users"
	"github.com/Rock2k3/notes-core/internal/infra/client"
	notesgrpcapi "github.com/Rock2k3/notes-grpc-api/v2/generated-sources"
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

	resp, err := usersGrpcClient.Client.GetUserByUUID(
		ctx,
		&notesgrpcapi.GetUserByUUIDRequest{UserUUID: uuid.String()},
	)
	if err != nil {
		return nil, err
	}

	user := &users.MyUser{
		UserUUID: uuid,
		Name:     resp.GetUser().Name,
	}

	return user, nil
}

func (a *usersGrpcAdapter) GetUserByName(name string) (*users.MyUser, error) {
	usersGrpcClient := client.NewUsersGrpcClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer usersGrpcClient.Conn.Close()
	defer cancel()

	resp, err := usersGrpcClient.Client.GetUserByName(
		ctx,
		&notesgrpcapi.GetUserByNameRequest{Name: name},
	)
	if err != nil {
		return nil, err
	}

	userUUID, _ := uuid.Parse(resp.GetUser().UserUUID)
	user := &users.MyUser{
		UserUUID: userUUID,
		Name:     name,
	}

	return user, nil
}

func (a *usersGrpcAdapter) AddUser(name string) (*users.MyUser, error) {
	usersGrpcClient := client.NewUsersGrpcClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer usersGrpcClient.Conn.Close()
	defer cancel()

	resp, err := usersGrpcClient.Client.AddUser(ctx, &notesgrpcapi.AddUserRequest{Name: name})
	if err != nil {
		return nil, err
	}

	userUUID, _ := uuid.Parse(resp.GetUser().UserUUID)
	user := &users.MyUser{
		UserUUID: userUUID,
		Name:     name,
	}

	return user, nil
}
