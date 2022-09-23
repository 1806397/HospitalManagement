package main

import (
	"HospitalManagement/Appointment/database"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAppointment(t *testing.T) {
	database.SetupConnection()
	recorder := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"patientName":  "Piyush K Singh",
		"patientPhone": 20973847934,
		"patientEmail": "Piyush@gmail.com",
		"doctorName":   "Neelesh Jain",
	}
	body, _ := json.Marshal(PostBody)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8050/setAppointment", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	SetAppointment(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
func TestChangeAppointment(t *testing.T) {
	database.SetupConnection()
	recorder := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"patientName":   "Piyush K Singh",
		"appointmentID": 22,
		"doctorName":    "Harsha K N",
		"patientPhone":  6264132730,
		"patientEmail":  "piyush.singh@gmail.com",
	}
	body, _ := json.Marshal(PostBody)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8050/changeAppoint", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	ChangeAppointment(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
func TestCancleAppointment(t *testing.T) {
	database.SetupConnection()
	recorder := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"id": 22,
	}
	body, _ := json.Marshal(PostBody)
	request, err := http.NewRequest(http.MethodDelete, "http://localhost:8050/cancelAppointment", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	CancelAppointment(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
