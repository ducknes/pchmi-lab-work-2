package database

import (
	_ "embed"
	"lab-work-2/domain"
)

//go:embed sql/get_all_employees.sql
var getAllEmployees string

func (r *LogisticRepository) GetAllEmployees() (employees []domain.Employee, err error) {
	if err = r.db.Select(&employees, getAllEmployees); err != nil {
		return
	}
	return
}

//go:embed sql/add_employee.sql
var addEmployee string

func (r LogisticRepository) AddEmployee(employee domain.Employee) (err error) {
	if _, err = r.db.NamedExec(addEmployee, employee); err != nil {
		return
	}
	return
}
