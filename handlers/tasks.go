package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func hello(c echo.Context) error  {
	return c.String(http.StatusOK, "Hello, World!")
}

func getTasks(c echo.Context) error  {
	return c.JSON(200, "GET Tasks")
}

func updateTask(c echo.Context) error  {
	return c.JSON(200, "PUT Tasks")
}

func deleteTask(c echo.Context) error  {
	return c.JSON(200, "DELETE Task " + c.Param("id"))
}

