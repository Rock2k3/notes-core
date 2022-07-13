package adapters

import (
	NotesGrpcApi "github.com/Rock2k3/notes-grpc-api/generated-sources"
	"github.com/google/uuid"
	"notes-core/internal/domain/users"
)

type usersGrpcAdapter struct {
}

func NewUsersGrpcAdapter() *usersGrpcAdapter {
	return &usersGrpcAdapter{}
}

func (a *usersGrpcAdapter) GetUserById(uuid uuid.UUID) (*users.MyUser, error) {
	usersGrpcClient := NewUsersGrpcClient()
	defer usersGrpcClient.conn.Close()
	defer usersGrpcClient.usersGrpcClientContext.cancel()

	resp, err := usersGrpcClient.client.GetUser(usersGrpcClient.usersGrpcClientContext.ctx, &NotesGrpcApi.GetUserRequest{UserId: uuid.String()})
	if err != nil {
		return nil, err
	}

	user := &users.MyUser{
		UserId: uuid,
		Name:   resp.GetUser().Name,
	}

	return user, nil
}

func (a *usersGrpcAdapter) AddUser(name string) (*users.MyUser, error) {
	usersGrpcClient := NewUsersGrpcClient()
	defer usersGrpcClient.conn.Close()
	defer usersGrpcClient.usersGrpcClientContext.cancel()

	resp, err := usersGrpcClient.client.AddUser(usersGrpcClient.usersGrpcClientContext.ctx, &NotesGrpcApi.AddUserRequest{Name: name})
	if err != nil {
		return nil, err
	}

	userId, _ := uuid.Parse(resp.GetUser().UserId)
	user := &users.MyUser{
		UserId: userId,
		Name:   name,
	}

	return user, nil
}
