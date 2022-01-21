package main

import (
	"fmt"
	"net/http"
	"nova/db"
	"strconv"

	"github.com/labstack/echo"
	j "github.com/ricardolonga/jsongo"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", func(c echo.Context) error {
		users, err := db.Users().AllG()
		justDie(err)
		return c.JSON(http.StatusOK, users)
	})
	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		user, err := db.Users(
			Where("id = ?", id),
		).OneG()
		justDie(err)
		return c.JSON(http.StatusOK, user)
	})
	e.GET("/users/search", func(c echo.Context) error {
		users, err := db.Users(
			Where("name = ?", c.QueryParam("name")),
		).AllG()
		fmt.Println("debug", users, err)
		justDie(err)
		return c.JSON(http.StatusOK, users)
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		user, err := db.FindUserG(id)
		justDie(err)
		rowsAff, err := user.DeleteG()
		fmt.Println("debug", rowsAff, err)
		json := j.Object().Put("rowsAff", rowsAff)
		justDie(err)
		return c.JSON(http.StatusOK, json)
	})
	e.POST("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	})
	e.PUT("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		user, err := db.Users(
			Where("id = ?", id),
		).OneG()
		justDie(err)
		user.Name = c.FormValue("name")
		return c.JSON(http.StatusOK, "")
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func justDie(err error) {
	if err != nil {
		panic(err)
	}
}
