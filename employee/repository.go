package employee

import (
	"database/sql"

	"github.com/GolangNortwindRestApi/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `select 
					id,first_name,
					last_name,company,
					email_address,job_title,
					business_phone,home_phone,
					COALESCE(mobile_phone,''),fax_number,address
				from employees
				limit ? offset ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var employees []*Employee
	for results.Next() {
		employee := &Employee{}
		err = results.Scan(&employee.ID, &employee.FirstName, &employee.Lastname, &employee.Company,
			&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
			&employee.MobilePhone, &employee.FaxNumber, &employee.Address)
		helper.Catch(err)
		employees = append(employees, employee)
	}
	return employees, nil
}

func (repo *repository) GetTotalEmployees() (int64, error) {
	const sql = "SELECT COUNT(*) FROM employees"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil

}

func (repo *repository) GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error) {
	const sql = `select 
					id,first_name,
					last_name,company,
					email_address,job_title,
					business_phone,home_phone,
					COALESCE(mobile_phone,''),fax_number,address
				from employees
				WHERE id=?`
	row := repo.db.QueryRow(sql, param.EmployeeId)
	employee := &Employee{}
	err := row.Scan(&employee.ID, &employee.FirstName, &employee.Lastname, &employee.Company,
		&employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone,
		&employee.MobilePhone, &employee.FaxNumber, &employee.Address)
	helper.Catch(err)

	return employee, nil
}
