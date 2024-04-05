package database

import (
	_ "embed"
	"lab-work-2/domain"
	"strconv"
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

func (r *LogisticRepository) AddEmployee(employee domain.Employee) (err error) {
	if _, err = r.db.NamedExec(addEmployee, employee); err != nil {
		return
	}
	return
}

//go:embed sql/delete_employee.sql
var deleteEmployee string

func (r *LogisticRepository) DeleteEmployee(id string) (err error) {
	idInt, _ := strconv.Atoi(id)
	if _, err = r.db.Exec(deleteEmployee, idInt); err != nil {
		return
	}
	return
}

//go:embed sql/get_employee.sql
var getEmployee string

func (r *LogisticRepository) GetEmployee(id string) (employee domain.Employee, err error) {
	idInt, _ := strconv.Atoi(id)
	if err = r.db.Get(&employee, getEmployee, idInt); err != nil {
		return
	}
	return
}

//go:embed sql/update_employee.sql
var updateEmployee string

func (r *LogisticRepository) UpdateEmployee(employee domain.Employee) (err error) {
	if _, err = r.db.NamedExec(updateEmployee, employee); err != nil {
		return
	}
	return
}
