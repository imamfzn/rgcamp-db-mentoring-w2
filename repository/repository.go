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
	rows := []string{}
	// tx := r.db.Raw("select name from employees where salary between ? AND ?", start, end).Scan(&rows)
	tx := r.db.Table("employees").Select("name").Where("salary BETWEEN ? AND ?", start, end).Scan(&rows)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

// select *
// from employees e
// join departments d on d.id = e.department_id
// WHERE d.name = 'Finance'
func (r *Repository) GetAllEmployeeByDepartmentName(department string) ([]model.Employee, error) {
	emps := []model.Employee{}
	tx := r.db.Table("employees e").
		Select("e.*").
		Joins("join departments d on d.id = e.department_id").
		Where("d.name = ?", department).
		Scan(&emps)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return emps, nil
}

func (r *Repository) GetAllEmployeeByRoomId(roomId int) ([]model.Employee, error) {
	emps := []model.Employee{}
	tx := r.db.Table("employees e").
		Select("e.*").
		Joins("join departments d on d.id = e.department_id").
		Where("d.room_id = ?", roomId).
		Scan(&emps)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return emps, nil
}

func (r *Repository) GetAverageSalaryByRoomLevel(level int) (float64, error) {
	return 0.0, nil
}

func (r *Repository) GetEmptyRooms() ([]model.Room, error) {
	return nil, nil
}
