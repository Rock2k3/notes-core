package routes

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userTO struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"Name"`
}

func AddUserRoutes(router *echo.Router) {
	router.Add(http.MethodGet, "/users/:uuid", handlerGetUserByUUID())
	router.Add(http.MethodPost, "/users", handlerAddUser())
}

func handlerGetUserByUUID() echo.HandlerFunc {

	return func(c echo.Context) error {
		userUUID, err := uuid.Parse(c.Param("uuid"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		//myUser := users.MyUser{UserId: userId}

		//myUser, err := users.GetUserByUUID(adapters.NewUsersGrpcAdapter(), userId)
		//if err != nil {
		//	if strings.Contains(err.Error(), "no rows in result set") {
		//		return c.String(http.StatusNotFound, fmt.Sprintf("Пользователя с id: %s не существует", userId))
		//	}
		//	return err
		//}

		user := userTO{
			UUID: userUUID,
			Name: "test name",
		}

		return c.JSON(http.StatusOK, user)
	}
}

func handlerAddUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		userUUID, _ := uuid.NewRandom()
		user := userTO{
			UUID: userUUID,
			Name: "test name",
		}

		return c.JSON(http.StatusOK, user)
	}
}
