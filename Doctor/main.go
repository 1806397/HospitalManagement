package main

import (
	"HospitalManagement/Doctor/PatientList"
	"HospitalManagement/Doctor/database"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func getPatient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var DocName PatientList.Doctorname
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &DocName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		PatientDetails, err := PatientList.Getmypatient(DocName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		PatientDetailsJSON, err := json.Marshal(PatientDetails)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(PatientDetailsJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}
}
func main() {
	http.HandleFunc("/myPatient", getPatient)
	database.SetupConnection()
	http.ListenAndServe(":9000", nil)
}
