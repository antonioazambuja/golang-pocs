package resources

import "errors"

// Employee - Employee
type Employee struct {
	Name      string `json:"name"`
	AccountID string `json:"accountId"`
}

var employees []*Employee

// NewEmployee - New employee
func NewEmployee() *Employee {
	return &Employee{
		Name:      "",
		AccountID: "",
	}
}

// GetEmployee - Get employee by name
func GetEmployee(employeeName string) (*Employee, error) {
	for _, employee := range employees {
		if employeeName == employee.Name {
			return employee, nil
		}
	}
	return NewEmployee(), errors.New("Not found employee")
}

// SaveEmployee - Save employee by name
func SaveEmployee(employeeName, accountID string) error {
	employees = append(employees, &Employee{
		Name:      employeeName,
		AccountID: accountID,
	})
	return nil
}
