package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Rock2k3/notes-core/internal/adapters"
	"github.com/Rock2k3/notes-core/internal/domain/users"
	"github.com/Rock2k3/notes-core/internal/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type userTO struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"Name"`
}

func RegisterUserRoutes(router *echo.Router) {
	log := logger.GetAppLogger()
	log.Debug("RegisterUserRoutes")

	usersGrpcAdapter := adapters.NewUsersGrpcAdapter()

	router.Add(http.MethodGet, "/users/:uuid", handlerGetUserByUUID(usersGrpcAdapter))
	router.Add(http.MethodPost, "/users", handlerAddUser(usersGrpcAdapter))
}

func handlerGetUserByUUID(users users.MyUsers) echo.HandlerFunc {
	//usersGrpcAdapter := adapters.NewUsersGrpcAdapter()

	return func(c echo.Context) error {
		userUUID, err := uuid.Parse(c.Param("uuid"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		//myUser, err := usersGrpcAdapter.GetUserByUUID(userUUID)
		myUser, err := users.GetUserByUUID(userUUID)
		if err != nil || myUser == nil {
			if strings.Contains(err.Error(), "no rows in result set") || myUser == nil {
				return c.String(http.StatusNotFound, fmt.Sprintf("Пользователя с UUID: %s не существует", userUUID))
			}
			return err
		}

		user := userTO{
			UUID: userUUID,
			Name: myUser.Name,
		}

		return c.JSON(http.StatusOK, user)
	}
}

func handlerAddUser(users users.MyUsers) echo.HandlerFunc {
	//usersGrpcAdapter := adapters.NewUsersGrpcAdapter()

	return func(c echo.Context) error {
		var u userTO
		err := json.NewDecoder(c.Request().Body).Decode(&u)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		//myUser, err := usersGrpcAdapter.GetUserByUUID(userUUID)
		myUser, err := users.AddUser(u.Name)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return c.String(http.StatusConflict, "Пользователь с таки именем уже существует")
			}
			return err
		}

		u.UUID = myUser.UserUUID

		return c.JSON(http.StatusOK, u)
	}
}
