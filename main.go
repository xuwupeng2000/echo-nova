package main

import (
	"fmt"
	"net/http"
	"nova/db"
	"strconv"

	"github.com/labstack/echo"
	j "github.com/ricardolonga/jsongo"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		user, err := db.FindUserG(id)
		justDie(err)
		return c.JSON(http.StatusOK, user)
	})
	e.GET("/users/search", func(c echo.Context) error {
		users, err := db.Users(
			Where("name = ?", c.QueryParam("name")),
		).AllG()
		justDie(err)
		return c.JSON(http.StatusOK, users)
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		user, err := db.FindUserG(id)
		justDie(err)
		rowsAff, err := user.DeleteG()
		json := j.Object().Put("rowsAff", rowsAff)
		justDie(err)
		return c.JSON(http.StatusOK, json)
	})
	e.POST("/users", func(c echo.Context) error {
		name := c.FormValue("name")
		age, _ := strconv.Atoi(c.FormValue("age"))
		user := new(db.User)
		user.Name = name
		user.Age = null.Int{age, true}
		user.InsertGP(boil.Infer())

		return c.JSON(http.StatusOK, user)
	})
	e.PUT("/users/:id", func(c echo.Context) error {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		user := db.FindUserGP(id)
		user.Name = c.FormValue("Name")
		age, _ := strconv.Atoi(c.FormValue("Age"))
		user.Age = null.Int{age, true}
		fmt.Println("debug: ", user)
		success := user.UpdateGP(boil.Infer())
		json := j.Object().Put("success", success).Put("user", user)

		return c.JSON(http.StatusOK, json)
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func justDie(err error) {
	if err != nil {
		panic(err)
	}
}
