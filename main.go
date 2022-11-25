package main

import (
	"mentoring/db"
	"mentoring/model"
)

func main() {
	credential := db.Credential{
		Host:         "db.fznlabs.xyz",
		Username:     "",
		Password:     "",
		DatabaseName: "",
		Port:         5432,
	}

	db, err := db.NewDB().Connect(credential)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(model.Department{}, model.Employee{}, model.Room{})
	if err != nil {
		panic(err)
	}
}
