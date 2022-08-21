package changeAppoint

type UpdateAppoint struct {
	Patient_name   string `json:"patientName"`
	Patient_phone  int    `json:"patientPhone"`
	Patient_email  string `json:"patientEmail"`
	Appointment_id int    `json:"appointmentID"`
	DocName        string `json:"doctorName"`
}
