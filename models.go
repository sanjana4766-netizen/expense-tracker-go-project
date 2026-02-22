package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ---------- MODELS ----------

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type Group struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type GroupMember struct {
	ID      uint `gorm:"primaryKey"`
	GroupID uint
	UserID  uint
}

type Expense struct {
	ID      uint `gorm:"primaryKey"`
	GroupID uint
	PaidBy  uint
	Amount  float64
}

type Split struct {
	ID        uint `gorm:"primaryKey"`
	ExpenseID uint
	UserID    uint
	Amount    float64
}

// ---------- DB INIT ----------

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("expense.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	DB.AutoMigrate(
		&User{},
		&Group{},
		&GroupMember{},
		&Expense{},
		&Split{},
	)
}
