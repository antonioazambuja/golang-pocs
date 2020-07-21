package services

import (
	rsc_v1 "github.com/antonioazambuja/golang-pocs/api-design/app/resources"
)

// GetEmployeeByName - Service get employee by name
func GetEmployeeByName(employeeName string) (*rsc_v1.Employee, error) {
	employee, errGetEmployee := rsc_v1.GetEmployee(employeeName)
	if errGetEmployee != nil {
		return nil, errGetEmployee
	}
	return employee, nil
}

// SaveEmployee - Service save employee by name
func SaveEmployee(employeeName, accountID string) error {
	errSaveEmployee := rsc_v1.SaveEmployee(employeeName, accountID)
	if errSaveEmployee != nil {
		return errSaveEmployee
	}
	return nil
}
