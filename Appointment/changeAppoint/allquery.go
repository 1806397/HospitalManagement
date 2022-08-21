package changeAppoint

import (
	"HospitalManagement/Appointment/database"
	"fmt"
)

func UpdatePatientAppoint(updateapp UpdateAppoint) error {
	var Id DocId
	result := database.Dbconn.QueryRow(`SELECT Doc_id FROM doctorlist WHERE Doc_name=?`, updateapp.DocName)
	result.Scan(&Id.Doc_id)
	fmt.Println(Id.Doc_id)
	_, err := database.Dbconn.Exec(`UPDATE appointment SET patient_name=?,patient_phone=?,patient_email=?,doctor_id=? WHERE appointment_id=?`, updateapp.Patient_name, updateapp.Patient_phone, updateapp.Patient_email, Id.Doc_id, updateapp.Appointment_id)
	if err != nil {
		return err
	}
	return nil
}
