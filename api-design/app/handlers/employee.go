package v1

import (
	"encoding/json"
	"log"
	"net/http"

	svc_v1 "github.com/antonioazambuja/golang-pocs/api-design/app/services"
	"github.com/gorilla/mux"
)

// HealthCheck - get summoner by name
func HealthCheck(w http.ResponseWriter, r *http.Request) {}

// GetEmployeeByName - Get employee by name
func GetEmployeeByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	employee, errGetEmployee := svc_v1.GetEmployeeByName(params["name"])
	if errGetEmployee != nil {
		log.Println(errGetEmployee)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employee)
	}
}

// SaveEmployee - Save employee
func SaveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	errGetEmployee := svc_v1.SaveEmployee(params["name"], params["accountID"])
	if errGetEmployee != nil {
		log.Println(errGetEmployee)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
