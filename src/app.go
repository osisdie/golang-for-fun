package main

import (
	"fmt"
	"net/http"
	"runtime"

	"example.com/app/printrange"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User
	c.BindJSON(&newUser)
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func main() {
	// Hello, Go go1.23.4
	fmt.Println("Hello, Go", runtime.Version())

	fmt.Println("Printing numbers with default separator:")
	printrange.PrintRangeNumbers(1, 5, "\n")

	fmt.Println("\nPrinting numbers with custom separator:")
	printrange.PrintRangeNumbers(2, 6, ", ")

	r := gin.Default()
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	r.GET("/users", getUsers)
	r.POST("/users", createUser)
	r.Run(":8080")
}
