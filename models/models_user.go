package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Register struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Username string `json:"username" validate:"required,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Profile_User struct {
	gorm.Model
	Employee_id string `json:"employee_id" validate:"required,min=3,max=32"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Birthday    string `json:"birthday"`
	Age         int    `json:"age"`
	IsActive    *bool  `json:"isactive" validate:"required"`
	Email       string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Tel         string `json:"tel"`
}
