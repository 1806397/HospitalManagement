package main

import (
	"HospitalManagement/Appointment/changeAppoint"
	"HospitalManagement/Appointment/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

func ChangeAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		err := CheckToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
func CheckToken(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return err
		}
		return err
	}
	tokenStr := cookie.Value
	claims := &changeAppoint.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}
		return err
	}
	if !tkn.Valid {
		return err
	}
	return nil

}
func CancelAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		err := CheckToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var cancel changeAppoint.Cancel
		bodybytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.Unmarshal(bodybytes, &cancel)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = changeAppoint.DeleteAppointment(cancel)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte(fmt.Sprintf("Appointment with id %d Successfully cancel", cancel.Appointment_id)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
func SetAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		err := CheckToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var app changeAppoint.Appoint
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
		AppointmentID, err := changeAppoint.InsertAppointment(app)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprintf("Appointment Created,your Appointment_id is %d", AppointmentID)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
func main() {
	http.HandleFunc("/changeAppoint", ChangeAppointment)
	http.HandleFunc("/cancelAppointment", CancelAppointment)
	http.HandleFunc("/setAppointment", SetAppointment)
	database.SetupConnection()
	http.ListenAndServe(":8050", nil)
}
