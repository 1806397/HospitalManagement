package PatientList

import (
	"HospitalManagement/Doctor/database"
)

func Getmypatient(doc Doctorname) ([]PatientDetails, error) {

	results, err := database.Dbconn.Query(`SELECT patient_name,patient_phone,patient_email FROM appointment WHERE doctor_id = (SELECT DOC_id FROM doctorlist WHERE Doc_name=?)`, doc.Name)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	patientslist := make([]PatientDetails, 0)
	for results.Next() {
		var patientlist PatientDetails
		results.Scan(&patientlist.Patient_name, &patientlist.Patient_phone, &patientlist.Patient_email)
		patientslist = append(patientslist, patientlist)
	}
	return patientslist, nil
}
