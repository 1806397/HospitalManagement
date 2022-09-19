package main

import (
	"HospitalManagement/Patient/database"
	"HospitalManagement/Patient/doctorList"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

var jwtKey = []byte("secret_key")

func DoctorList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := CheckToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		doctorName, err := doctorList.GetDoctor()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		doctorJSON, err := json.Marshal(doctorName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(doctorJSON)
	case http.MethodPost:
		err := CheckToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var DoctorName doctorList.DoctorDB
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &DoctorName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		doctorData, err := doctorList.GetDoctorByTimeSlot(DoctorName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		doctorDataJSON, err := json.Marshal(doctorData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(doctorDataJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
func CheckToken(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return err
		}
		return err
	}
	tokenStr := cookie.Value
	claims := &doctorList.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}
		return err
	}
	if !tkn.Valid {
		return err
	}
	return nil

}
func SetAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		err := CheckToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var app doctorList.Appoint
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &app)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		AppointmentID, err := doctorList.InsertAppointment(app)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprintf("Appointment Created,your Appointment_id is %d", AppointmentID)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func CreateRegistration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var reg doctorList.Registration
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &reg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = doctorList.RegistrationUser(reg)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		w.Write([]byte(fmt.Sprintf("User with ID %s Successfully created", reg.User_id)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
func UserLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var login doctorList.Login
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &login)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = doctorList.UserLogin(login)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// w.Write([]byte(fmt.Sprintf("User %s Logged in Successfully", login.User_id)))
		//Token creation
		expirationTime := time.Now().Add(time.Minute * 5)
		claims := &doctorList.Claims{
			Username: login.User_id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

/*func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
	tokenStr := cookie.Value
	claims := &doctorList.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// w.Write([]byte(fmt.Sprintf("Welcome user,%s", claims.Username)))
}*/
func main() {
	http.HandleFunc("/getDoctor", DoctorList)
	http.HandleFunc("/setAppointment", SetAppointment)
	http.HandleFunc("/registration", CreateRegistration)
	http.HandleFunc("/login", UserLogin)

	// http.HandleFunc("/home", Home)
	database.SetupConnection()
	http.ListenAndServe(":8000", nil)
}
