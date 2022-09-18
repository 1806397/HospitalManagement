package main

import (
	"HospitalManagement/Appointment/changeAppoint"
	"HospitalManagement/Appointment/database"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ChangeAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var appoint changeAppoint.UpdateAppoint
		bodybytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodybytes, &appoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = changeAppoint.UpdatePatientAppoint(appoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{Message:"Appointment Changed"}`))
	}
}
func main() {
	http.HandleFunc("/changeAppoint", ChangeAppointment)
	database.SetupConnection()
	http.ListenAndServe(":8050", nil)
}
