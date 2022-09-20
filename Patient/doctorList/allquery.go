package doctorList

import (
	"HospitalManagement/Patient/database"
	"errors"
)

// UserExists=errors.New("User ID already exist")
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

func RegistrationUser(reg Registration) error {
	UserExist := errors.New("user already exist")
	PasswordMismatch := errors.New("password Mismatch.Please check your password")
	_, err := database.Dbconn.Exec(`INSERT INTO registration (UserID,Password) VALUES (?,?)`, reg.User_id, reg.Password)
	if err != nil {
		// fmt.Println("line 79")
		return UserExist
	}
	if reg.Password != reg.ConfirmPassword {
		return PasswordMismatch
	}
	return nil
}

func UserLogin(log Login) error {
	var UserLog Login
	results := database.Dbconn.QueryRow(`SELECT * FROM registration WHERE UserID=?`, log.User_id)
	err := results.Scan(&UserLog.User_id, &UserLog.Password)
	if err != nil {
		return errors.New("user doesn't exist")
	}
	if log.Password != UserLog.Password && log.User_id == UserLog.User_id {
		return errors.New("incorrect Username or Password")
	}
	return nil

}

//select *from getLastRecord ORDER BY id DESC LIMIT 1
