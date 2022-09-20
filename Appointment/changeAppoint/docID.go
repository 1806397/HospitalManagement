package changeAppoint

import "github.com/golang-jwt/jwt"

type DocId struct {
	Doc_id int `json:"DocID"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type UpdateAppoint struct {
	Patient_name   string `json:"patientName"`
	Patient_phone  int    `json:"patientPhone"`
	Patient_email  string `json:"patientEmail"`
	Appointment_id int    `json:"appointmentID"`
	DocName        string `json:"doctorName"`
}
type Cancel struct {
	Appointment_id int `json:"id"`
}
type Appoint struct {
	Patient_name  string `json:"patientName"`
	Patient_phone int    `json:"patientPhone"`
	Patient_email string `json:"patientEmail"`
	Doctor_name   string `json:"doctorName"`
}
type DoctorDB struct {
	Doc_id   int    `json:"DocID"`
	Doc_name string `json:"name"`
	In_time  string `json:"In"`
	Out_time string `json:"Out"`
}
