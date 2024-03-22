package main

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLoginRequest struct {
	Admin    string `json:"admin"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Id        int    `json:"id"`
	Username  string `json:"username"`
	SessionId string `json:"sessionid"`
}

type AdminLoginResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Id        int    `json:"id"`
	Admin     string `json:"admin"`
	Username  string `json:"username"`
	SessionId string `json:"sessionid"`
}

type CheckSession struct {
	Username  string `json:"username"`
	SessionId string `json:"sessionid"`
}

type CheckSessionRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Id      int    `json:"id"`
}

type StandardIdRequest struct {
	Id int `json:"id"`
}

type StandardResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    ProfileInfo `json:"data"`
}

type ProfileInfo struct {
	Patient       Patient         `json:"patient"`
	AllergiesInfo []AllergiesInfo `json:"allergiesInfo"`
	ComplaintInfo []ComplaintInfo `json:"complaintInfo"`
	DoctorNotes   []DoctorNotes   `json:"doctorNotes"`
	FamilyHistory []FamilyHistory `json:"familyHistory"`
	Medications   []Medications   `json:"medications"`
	UpcomingDates []UpcomingDates `json:"upcomingDates"`
	Uploads       []Uploads       `json:"uploads"`
	Vaccine       []Vaccine       `json:"vaccine"`
}

type Patient struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Sex             string `json:"sex"`
	Address         string `json:"address"`
	DOB             string `json:"dob"`
	BType           string `json:"btype"`
	Allergies       string `json:"allergies"`
	Reactions       string `json:"reactions"`
	CMC             string `json:"cmc"`
	FamilyCMC       string `json:"familycmc"`
	MentalCondition string `json:"mentalcondition"`
	Surgeries       string `json:"surgeries"`
	Complaint       string `json:"complaint"`
	Marital         string `json:"marital"`
	Smoke           string `json:"smoke"`
	Contacts        string `json:"contacts"`
	Number          string `json:"number"`
	Height          int    `json:"height"`
	Weight          int    `json:"weight"`
	IceInstructions string `json:"ice_instructions"`
	Lang            string `json:"lang"`
	Role            string `json:"role"`
	Class           string `json:"class"`
	ProfileLink     string `json:"profile_link"`
	Profile         []byte `json:"profile"`
}

type AllergiesInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Reaction  string `json:"reaction"`
	PatientId int    `json:"patientId"`
}

type ComplaintInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Descriprion string `json:"description"`
	Date        string `json:"date"`
	PatientId   int    `json:"patientId"`
}

type DoctorNotes struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Number     string `json:"number"`
	Location   string `json:"location"`
	Notes      string `json:"notes"`
	Speciality string `json:"speciality"`
	PatientId  int    `json:"patientId"`
}

type FamilyHistory struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Relationship      string `json:"relationship"`
	MedicalConditions string `json:"medicalConditions"`
	PatientId         int    `json:"patientId"`
}

type Medications struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
	PatientId int    `json:"patientId"`
}

type UpcomingDates struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	PatientId int    `json:"patientId"`
}

type Uploads struct {
	Id        int    `json:"id"`
	File      []byte `json:"file"`
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	PatientId int    `json:"patientId"`
}
type Vaccine struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Details   string `json:"details"`
	Date      string `json:"date"`
	PatientId int    `json:"patientId"`
}

type BasicInfoStruct struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	DOB             string `json:"dob"`
	Btype           string `json:"btype"`
	Address         string `json:"address"`
	Gender          string `json:"gender"`
	MartialStatus   string `json:"martialStatus"`
	Allergies       string `json:"allergies"`
	Height          string `json:"height"`
	Weight          string `json:"weight"`
	Smoke           string `json:"smoke"`
	Surgeries       string `json:"surgeries"`
	Contacts        string `json:"contacts"`
	CMC             string `json:"cmc"`
	MentalCon       string `json:"mentalCon"`
	IceInstructions string `json:"iceInstructions"`
	Profile         []byte `json:"profile"`
	ProfileLink     string `json:"profile_link"`
}

type EditCurrentComplaintStruct struct {
	Id        int    `json:"id"`
	Complaint string `json:"complaint"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	DOB      string `json:"dob"`
	Address  string `json:"address"`
	PhNumber string `json:"contacts"`
	Gender   string `json:"sex"`
	BType    string `json:"btype"`
}

type DownFile struct {
	FileID    int `json:"id"`
	PatientId int `json:"patientId"`
}

type DownFileREs struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	File    []byte `json:"file"`
}

type ProfilePic struct {
	PicID int `json:"id"`
}

type ProfilePicRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Picture []byte `json:"picture"`
}
