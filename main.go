package main

import (
	"fmt"

	"mentoring/db"
	"mentoring/model"
	"mentoring/repository"
)

func printEmployees(emps []model.Employee) {
	for _, emp := range emps {
		fmt.Println(emp.ID, emp.Name)
	}
}

func main() {
	credential := db.Credential{
		Host:         "",
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

	repo := repository.NewRepo(db)

	start := 7000000
	end := 10000000
	names, err := repo.GetAllEmployeeNameBySalaryRange(start, end)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Employee dengan range gaji %d - %d: %+v\n", start, end, names)

	departmentName := "Research & Development"
	emps, err := repo.GetAllEmployeeByDepartmentName(departmentName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Employee di Department %s:\n", departmentName)
	printEmployees(emps)

	roomId := 1
	emps, err = repo.GetAllEmployeeByRoomId(roomId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Employee yang kerja di ruangan: %d:\n", roomId)
	printEmployees(emps)

	roomLevel := 10
	avg, err := repo.GetAverageSalaryByRoomLevel(roomLevel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Rata-rata salary di lantai %d: %f\n", roomLevel, avg)

	rooms, err := repo.GetEmptyRooms()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ruangan tanpa penghuni:", rooms)
}
