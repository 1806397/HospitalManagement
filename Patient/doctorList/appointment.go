package doctorList

import (
	jwt "github.com/dgrijalva/jwt-go"
)

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
type DoctorDB struct {
	Doc_id   int    `json:"DocID"`
	Doc_name string `json:"name"`
	In_time  string `json:"In"`
	Out_time string `json:"Out"`
}
