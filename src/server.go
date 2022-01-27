package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	router := NewRouter()

	router.Logger.Fatal(router.Start(":8080"))
}

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", userIndexHandler)
	e.GET("/users/:name", userShowHandler)
	e.POST("users", userCreateHandler)

	return e
}

//func helloHandler(c echo.Context) error {
//	return c.String(http.StatusOK, "Hello, Echo World!!")
//}

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email""`
}

type Users []User

func userIndexHandler(c echo.Context) error {
	var users Users

	users = append(users, User{
		Name:  "Taro",
		Email: "taro@example.com",
	})
	users = append(users, User{
		Name:  "Jiro",
		Email: "jiro@example.com",
	})

	return c.JSON(http.StatusOK, users)
}

func userShowHandler(c echo.Context) error {
	var user User
	name := c.Param("name")

	if name == "taro" {
		user = User{
			Name:  "Taro",
			Email: "taro@example.com",
		}
	} else if name == "jiro" {
		user = User{
			Name:  "Jiro",
			Email: "jiro@example.com",
		}
	}

	return c.JSON(http.StatusOK, user)
}

func userCreateHandler(c echo.Context) (err error) {
	user := new(User)

	if err = c.Bind(user); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, user)
}
