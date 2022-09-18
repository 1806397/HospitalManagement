package doctorList

import (
	"HospitalManagement/Patient/database"
	"fmt"
)

func GetDoctor() ([]DoctorDB, error) {
	results, err := database.Dbconn.Query(`SELECT Doc_id,Doc_name,In_time,Out_time FROM doctorlist`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	doctors := make([]DoctorDB, 0)
	for results.Next() {
		var doctor DoctorDB
		results.Scan(&doctor.Doc_id, &doctor.Doc_name, &doctor.In_time, &doctor.Out_time)
		doctors = append(doctors, doctor)
	}
	return doctors, nil

}
func GetDoctorByTimeSlot(doctor DoctorDB) ([]DoctorDB, error) {
	results, err := database.Dbconn.Query(`SELECT Doc_id,Doc_name,In_time,Out_time FROM doctorlist WHERE In_time<=? AND Out_time>=?`, doctor.In_time, doctor.Out_time)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	doctors := make([]DoctorDB, 0)
	for results.Next() {
		var doctorSlot DoctorDB
		results.Scan(&doctorSlot.Doc_id, &doctorSlot.Doc_name, &doctorSlot.In_time, &doctorSlot.Out_time)
		doctors = append(doctors, doctorSlot)
	}
	return doctors, nil
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

//select *from getLastRecord ORDER BY id DESC LIMIT 1
