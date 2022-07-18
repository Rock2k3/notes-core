package adapters

import (
	"errors"
	"fmt"
	"github.com/Rock2k3/notes-core/internal/domain/users"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
)

const baseUsersPath = "build"

type userFilesAdapter struct {
}

func NewUserFilesAdapter() *userFilesAdapter {
	return &userFilesAdapter{}
}

func (a *userFilesAdapter) AddUser(name string) (*users.MyUser, error) {
	path := userIsExists(name, uuid.Nil)
	if path != "" {
		return nil, errors.New("пользователь с таки именем уже существует")
	}

	userId, _ := createUser(name)
	return &users.MyUser{
		UserId: userId,
		Name:   name,
	}, nil
}

func (a *userFilesAdapter) GetUserById(uuid uuid.UUID) (*users.MyUser, error) {
	return &users.MyUser{
		UserId: uuid,
		Name:   "",
	}, nil
}

func GetUserByName(name string) users.MyUser {
	fileName := userIsExists(name, uuid.Nil)

	if fileName != "" {
		userId, _ := uuid.Parse(strings.Split(fileName, "_")[0])
		return users.MyUser{
			UserId: userId,
			Name:   name,
		}
	} else {
		userId, err := createUser(name)
		if err != nil {
			return users.MyUser{}
		}
		return users.MyUser{UserId: userId, Name: name}
	}
}

func createUser(name string) (uuid.UUID, error) {
	userId, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	filename := fmt.Sprintf("build/%s_%s", userId.String(), name)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	return userId, nil
}

func findUserByName(name string) (uuid.UUID, error) {
	filename := fmt.Sprintf("build/%s", name)

	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	userId, err := uuid.FromBytes(f)
	if err != nil {
		panic(err)
	}

	return userId, nil
}

func userIsExists(name string, userId uuid.UUID) string {
	var filePath string
	if name != "" {
		filePathArr, err := filepath.Glob(fmt.Sprintf("%s/*_%s", baseUsersPath, name))
		if err != nil || len(filePathArr) < 1 {
			return ""
		}
		filePath = filePathArr[0]
	} else if userId != uuid.Nil {
		filePathArr, err := filepath.Glob(fmt.Sprintf("%s/%s_*", baseUsersPath, name))
		if err != nil || len(filePathArr) < 1 {
			return ""
		}
		filePath = filePathArr[0]
	}

	_, err := os.Stat(filePath)
	if err != nil {
		return ""
	}

	str := strings.Split(filePath, baseUsersPath)[1]

	return str[1 : len(str)-1]
}
