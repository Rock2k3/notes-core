package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"notes-core/internal/adapters"
	"notes-core/internal/domain/users"
	"notes-core/pkg/notesHttpServer"
	"strings"
)

type userTO struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"Name"`
}

func GetUserRoutes() []notesHttpServer.Route {
	var r []notesHttpServer.Route

	r = append(r, notesHttpServer.Route{
		Method:  http.MethodGet,
		Path:    "/users/:id",
		Handler: getUser,
	})
	r = append(r, notesHttpServer.Route{
		Method:  http.MethodPost,
		Path:    "/users",
		Handler: createUser,
	})

	return r
}

func createUser(c echo.Context) error {
	var u userTO
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	myUser, err := users.AddUser(adapters.NewUsersGrpcAdapter(), u.Name)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return c.String(http.StatusConflict, "Пользователь с таки именем уже существует")
		}
		return err
	}
	fmt.Println("userGrpcAdapter", myUser)

	return c.JSON(http.StatusOK, myUser)
}

func getUser(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	myUser, err := users.GetUserById(adapters.NewUsersGrpcAdapter(), userId)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return c.String(http.StatusNotFound, fmt.Sprintf("Пользователя с id: %s не существует", userId))
		}
		return err
	}

	return c.JSON(http.StatusOK, myUser)
}
