package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//  go get -u github.com/gin-gonic/gin
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {

	newServer().Run()

}

func newServer() *gin.Engine {
	r := gin.Default()

	r.GET("", helloHandler)

	r.GET("/:name", helloUserHandler)

	r.POST("/add", helloAccoundHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}

func helloHandler(c *gin.Context) {

	db, err := sql.Open("mysql", "admin:examboomadmin@tcp(examboom.calza1qfpu2i.ap-northeast-2.rds.amazonaws.com:3306)/examboom")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var email string
	err = db.QueryRow("select email from User where nickname = '한승수'").Scan(&email)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(email)

	c.JSON(http.StatusOK, gin.H{
		"responseData": "hello world : " + email,
	})
}

func helloUserHandler(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"greetings": fmt.Sprintf("hello %v", name),
	})
}

func helloAccoundHandler(c *gin.Context) {
	var data Account
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("err: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"dataReceived": data,
	})
}
