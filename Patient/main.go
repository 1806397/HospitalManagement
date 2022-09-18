package main

import (
	"HospitalManagement/Patient/database"
	"HospitalManagement/Patient/doctorList"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func DoctorList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doctorName, err := doctorList.GetDoctor()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		doctorJSON, err := json.Marshal(doctorName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(doctorJSON)
	case http.MethodPost:
		var DoctorName doctorList.DoctorDB
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &DoctorName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		doctorData, err := doctorList.GetDoctorByTimeSlot(DoctorName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		doctorDataJSON, err := json.Marshal(doctorData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(doctorDataJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
func SetAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var app doctorList.Appoint
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &app)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		AppointmentID, err := doctorList.InsertAppointment(app)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprintf("Appointment Created,your Appointment_id is %d", AppointmentID)))

	}
}
func main() {
	http.HandleFunc("/getDoctor", DoctorList)
	http.HandleFunc("/setAppointment", SetAppointment)
	database.SetupConnection()
	http.ListenAndServe(":8000", nil)
}
