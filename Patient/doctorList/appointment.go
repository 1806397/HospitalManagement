package doctorList

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Appoint struct {
	Patient_name  string `json:"patientName"`
	Patient_phone int    `json:"patientPhone"`
	Patient_email string `json:"patientEmail"`
	Doctor_name   string `json:"doctorName"`
}

type Registration struct {
	User_id         string `json:"userID"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"confirmPassword"`
}
type Login struct {
	User_id  string `json:"userID"`
	Password string `json:"Password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type Cancel struct {
	Appointment_id int `json:"id"`
}
