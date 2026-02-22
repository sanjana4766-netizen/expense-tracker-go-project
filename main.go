package main

import "github.com/gin-gonic/gin"

func main() {
	InitDB()
	r := gin.Default()

	r.POST("/users", CreateUser)
	r.POST("/groups", CreateGroup)
	r.POST("/groups/:id/add", AddUserToGroup)
	r.POST("/expenses", AddExpense)
	r.GET("/settlement/:groupId", GetSettlement)

	r.Run(":8080")
}