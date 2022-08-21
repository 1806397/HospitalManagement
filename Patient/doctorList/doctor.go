package doctorList

type DoctorDB struct {
	Doc_id   int    `json:"DocID"`
	Doc_name string `json:"name"`
	In_time  string `json:"In"`
	Out_time string `json:"Out"`
}
