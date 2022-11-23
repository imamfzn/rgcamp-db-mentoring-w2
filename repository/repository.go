package repository

import (
	"mentoring/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllEmployeeNameBySalaryRange(start, end int) ([]string, error) {
	return nil, nil
}

func (r *Repository) GetAllEmployeeByDepartmentName(department string) ([]model.Employee, error) {
	return nil, nil
}

func (r *Repository) GetAllEmployeeByRoomId(roomId int) ([]model.Employee, error) {
	return nil, nil
}

func (r *Repository) GetAverageSalaryByRoomLevel(level int) (float64, error) {
	return 0.0, nil
}

func (r *Repository) GetEmptyRooms() ([]model.Room, error) {
	return nil, nil
}
