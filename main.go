package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mzc-devops-toyproject/gateway-service/models"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var logPath = path.Dir("/var/log/")

func main() {
	logFile, err := os.OpenFile(logPath+"/moodi-gw.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		PrintToJSON(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	port := flag.Int("port", 80, "Port")

	if !flag.Parsed() {
		flag.Parse()
	}

	e := echo.New()
	logConfig := &middleware.LoggerConfig{
		Output: logFile,
	}
	e.Use(middleware.LoggerWithConfig(*logConfig))
	e.Use(middleware.Recover())
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
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, `<div style="text-align: center;">
		<h1>Welcome to Moodi</h1>
		<img src="./sunny.svg" />
		</div>`)
	})
	e.GET("/sunny.svg", func(c echo.Context) error {
		if err != nil {
			return err
		}
		return c.File("./public/sunny.png")
	})
	e.GET(`health-check`, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, models.ResponseJSON{
			RequestID: bson.NewObjectId(),
			Message:   `Alive`,
			Code:      200,
			Timestamp: time.Now().Unix(),
			Data:      ``,
		})
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*port)))
}

func PrintToJSON(obj interface{}) {
	s, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(s))
}
