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
func DeleteAppointment(cancel Cancel) error {
	_, err := database.Dbconn.Exec(`DELETE FROM appointment WHERE appointment_id=?`, cancel.Appointment_id)
	if err != nil {
		return err
	}
	return nil

}
func InsertAppointment(app Appoint) (int, error) {
	var doctor DoctorDB
	results := database.Dbconn.QueryRow(`SELECT Doc_id FROM doctorlist WHERE Doc_name=?`, app.Doctor_name)
	err := results.Scan(&doctor.Doc_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	_, err = database.Dbconn.Exec(`INSERT INTO appointment (patient_name,patient_phone,patient_email,doctor_id) VALUES (?,?,?,?)`, app.Patient_name, app.Patient_phone, app.Patient_email, doctor.Doc_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	results = database.Dbconn.QueryRow(`SELECT appointment_id FROM appointment ORDER BY appointment_id DESC LIMIT 1`)
	err = results.Scan(&doctor.Doc_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return doctor.Doc_id, nil
}
