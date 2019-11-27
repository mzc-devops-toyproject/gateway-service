package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/mzc-devops-toyproject/gateway-service/models"
	"flag"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.Int("port", 80, "Port")

	if !flag.Parsed() {
		flag.Parse()
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Setup proxy
	// url1, err := url.Parse("http://localhost:8081")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// url2, err := url.Parse("http://localhost:8082")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// targets := []*middleware.ProxyTarget{
	// 	{
	// 		URL: url1,
	// 	},
	// 	{
	// 		URL: url2,
	// 	},
	// }
	// e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<div style="text-align: center;">
		<h1>Welcome to Moodi</h1>
		<img src="./sunny.svg" />
		</div>`)
	})
	e.GET("/sunny.svg", func(c echo.Context) error {
		return c.File("./public/sunny.png")
	})
	e.GET(`health-check`, func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.ResponseJSON{
			RequestID: bson.NewObjectId(),
			Message: `Alive`,
			Code: 200,
			Timestamp: time.Now().Unix(),
			Data: ``,
		})
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*port)))
}
