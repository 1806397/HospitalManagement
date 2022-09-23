package main

import (
	"HospitalManagement/Patient/database"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*func CreateToken(w http.ResponseWriter, User_id string) {
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &doctorList.Claims{
		Username: User_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}*/
func TestDoctorList(t *testing.T) {
	database.SetupConnection()
	recorder := httptest.NewRecorder()
	// http.SetCookie(recorder, &http.Cookie{Name: "token", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlBpeXVzaFMiLCJleHAiOjE2NjM4MzM0NTJ9.NTE-zlEDnY_SlFBFmfO2xz0u3iKTMvV11DcQwKJKFe0", Expires: time.Now().Add(5 * time.Minute)})
	PostBody := map[string]interface{}{
		"In":  "7",
		"Out": "9",
	}
	body, _ := json.Marshal(PostBody)
	request, err := http.NewRequest(http.MethodPost, "localhost:8000/registration", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	DoctorList(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
func TestCreateRegistration(t *testing.T) {
	database.SetupConnection()
	PostBody := map[string]interface{}{
		"userID":          "PiyushS",
		"Password":        "1234",
		"confirmPassword": "1234",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "localhost:8000/registration", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	CreateRegistration(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
func TestUserLogin(t *testing.T) {
	database.SetupConnection()
	PostBody := map[string]interface{}{
		"userID":   "PiyushS",
		"Password": "Piyush12",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "localhost:8000/login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	UserLogin(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
