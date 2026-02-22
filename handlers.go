package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	DB.Create(&u)
	c.JSON(http.StatusOK, u)
}

func CreateGroup(c *gin.Context) {
	var g Group
	c.BindJSON(&g)
	DB.Create(&g)
	c.JSON(http.StatusOK, g)
}

func AddUserToGroup(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Param("id"))
	var req struct{ UserID uint }
	c.BindJSON(&req)

	m := GroupMember{GroupID: uint(gid), UserID: req.UserID}
	DB.Create(&m)
	c.JSON(http.StatusOK, m)
}

func AddExpense(c *gin.Context) {
	var req struct {
		GroupID uint
		PaidBy  uint
		Amount  float64
	}
	c.BindJSON(&req)

	e := Expense{GroupID: req.GroupID, PaidBy: req.PaidBy, Amount: req.Amount}
	DB.Create(&e)

	var members []GroupMember
	DB.Where("group_id = ?", req.GroupID).Find(&members)

	split := req.Amount / float64(len(members))
	for _, m := range members {
		DB.Create(&Split{ExpenseID: e.ID, UserID: m.UserID, Amount: split})
	}

	c.JSON(http.StatusOK, e)
}

func GetSettlement(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Param("groupId"))
	c.JSON(http.StatusOK, CalculateSettlement(uint(gid)))
}