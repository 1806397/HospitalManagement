package PatientList

type PatientDetails struct {
	Patient_name  string `json:"patientName"`
	Patient_phone int    `json:"patientPhone"`
	Patient_email string `json:"patientEmail"`
	// Doctor_name   string `json:"doctorName"`
}
