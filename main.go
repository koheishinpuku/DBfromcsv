package main

import (
	// "fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"local.packages/Controllers"
	// "time"
	// "os"
	// "local.packages/DB"
	// "github.com/gomodule/redigo/redis"
	// "bufio"
)

// type Table1 struct {
//   Id       int    `json:id sql:AUTO_INCREMENT`
//   Name     string  `json:name`
//   Item     string `json:item`
// }

// func show(c echo.Context) error {
// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")
// 	return c.String(http.StatusOK, "team:"+team+", member:"+member)
// }
// func Handler_hello(c echo.Context) error {
// 	name := c.FormValue("name")
// 	return c.String(http.StatusOK, "create name = "+name)
// }

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("operation/extract", Controllers.CreateFile())
	e.POST("operation/import", Controllers.ImportData())
	// e.GET("/items", show)
	// e.POST("/hello", Handler_hello)
	e.Logger.Fatal(e.Start(":5000"))

}

// db := DB.GormConnect()
// db.CreateTable(&Table1{})
//
// eventEx := Table1{}
// CreateUser := Table1{}
// CreateUser.Name = "ケンちゃん"
// CreateUser.Item = "カステラ"
// db.Create(&CreateUser)
// db.First(&eventEx, "id = ?", 3)
// fmt.Println(eventEx)
