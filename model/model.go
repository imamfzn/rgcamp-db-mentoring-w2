package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	Name         string `gorm:"type:varchar(100);not null"`
	Age          int    `gorm:"type:int;not null"`
	Salary       int
	Address      string `gorm:"type:varchar(255)"`
	DepartmentId int
}

type Department struct {
	gorm.Model

	Name   string
	RoomId int
}

type Room struct {
	gorm.Model

	Code  string `gorm:"type:varchar(10);unique_index"`
	Level int    `gorm:"not null"`
}

type EmployeeRoom struct {
	Name  string
	Code  string
	Level int
}
