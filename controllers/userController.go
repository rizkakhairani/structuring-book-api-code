package controllers

import (
	"book-api/lib/database"
	"book-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 	"success",
		"users":	users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	user, e := database.GetUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":	"success get user by id",
		"user":		user,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	user, e := database.CreateUser(user)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":	"success create new user",
		"user":		user,
	})
}

func UpdateUserController(c echo.Context) error {
	var newUser models.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	user, e := database.UpdateUser(id, newUser)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":	"success update user by id",
		"user":		user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	user, e := database.DeleteUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":	"success delete user by id",
		"user":		user,
	})
}
