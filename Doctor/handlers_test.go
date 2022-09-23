package main

import (
	"HospitalManagement/Doctor/database"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPatient(t *testing.T) {
	database.SetupConnection()
	recorder := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"name": "Harsha K N",
	}
	body, _ := json.Marshal(PostBody)
	request, err := http.NewRequest(http.MethodPost, "localhost:9000/myPatient", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	getPatient(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
