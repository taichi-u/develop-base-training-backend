package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Profile struct {
	Name      string `json:"name,omitempty"`       //パスカルケースからキャラメルケースへ
	Message   string `json:"message,omitempty"`    //パスカルケースからキャラメルケースへ
	Date      int    `json:"date,omitempty"`       //パスカルケースからキャラメルケースへ
	Boolcheck bool   `json:"bool_check,omitempty"` //パスカルケースからスネークケースへ
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	// 自己紹介エンドポイント
	e.GET("/profile", func(c echo.Context) error {
		profile := Profile{
			Name:      "taichi",
			Message:   "私はGoが大好きなエンジニアです！よろしくお願いします。",
			Date:      309,
			Boolcheck: true,
		}
		return c.JSON(http.StatusOK, profile)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
