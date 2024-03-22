package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initializing() {
	username := "root"
	password := "golangMarwan"
	ip := "localhost"
	port := "3306"
	dbname := "fiephrs"

	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp("+ip+":"+port+")/"+dbname)
	if err != nil {
		println(err.Error())
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var (
		logReq LoginRequest
		logRes LoginResponse
	)
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		logRes.Success = false
		logRes.Message = err.Error()
		json.NewEncoder(w).Encode(logRes)
		return
	}

	checkValidSessions(logReq.Username)
	validLogin := checkUser(logReq.Username, logReq.Password)
	if validLogin {
		sessionid, timeLeft := generateSessionAndTimeLeft(logReq.Username)
		dbInsertSession(logReq.Username, sessionid, timeLeft)
		logRes.Success = true
		logRes.Message = "Login Successfully"
		logRes.Username = logReq.Username
		logRes.SessionId = string(sessionid[:])
		logRes.Id, err = dbGetUserId(logReq.Username)
		if err != nil {
			logRes.Success = false
			logRes.Message = err.Error()
			json.NewEncoder(w).Encode(logRes)
			return
		}
	} else {
		logRes.Success = false
		logRes.Message = "Login Failed"
	}
	json.NewEncoder(w).Encode(logRes)
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	var (
		logReq AdminLoginRequest
		logRes AdminLoginResponse
	)
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		logRes.Success = false
		logRes.Message = err.Error()
		json.NewEncoder(w).Encode(logRes)
		return
	}

	checkValidSessions(logReq.Admin)
	validLogin := checkAdmin(logReq.Admin, logReq.Password)
	if validLogin {
		sessionid, timeLeft := generateSessionAndTimeLeft(logReq.Admin)
		dbInsertSession(logReq.Admin, sessionid, timeLeft)
		logRes.Success = true
		logRes.Message = "Login Successfully"
		logRes.Admin = logReq.Admin
		logRes.SessionId = string(sessionid[:])
		logRes.Id, err = dbGetAdminId(logReq.Admin)
		logRes.Username = logReq.Username
		if err != nil {
			logRes.Success = false
			logRes.Message = err.Error()
			json.NewEncoder(w).Encode(logRes)
			return
		}
	} else {
		logRes.Success = false
		logRes.Message = "Login Failed"
	}
	json.NewEncoder(w).Encode(logRes)
}

func checkSession(w http.ResponseWriter, r *http.Request) {
	var (
		request  CheckSession
		response CheckSessionRes
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	username, err := dbCheckSession(request.SessionId)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	Id, err := dbGetUserId(username)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Session is valid!"
	response.Id = Id
	json.NewEncoder(w).Encode(response)
}

func checkAdminSession(w http.ResponseWriter, r *http.Request) {
	var (
		request  CheckSession
		response CheckSessionRes
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	username, err := dbCheckSession(request.SessionId)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	Id, err := dbGetAdminId(username)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Session is valid!"
	response.Id = Id
	json.NewEncoder(w).Encode(response)
}

func getProfileInfo(w http.ResponseWriter, r *http.Request) {
	var (
		request     StandardIdRequest
		response    StandardResponse
		profileInfo ProfileInfo
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	patient, err := dbGetPatientInfo(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	allergies, err := dbGetAllergies(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	complaint, err := dbGetComplaint(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	doctorNotes, err := dbGetDoctorNotes(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	familyHistory, err := dbGetFamilyHistory(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	medications, err := dbGetMedications(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	upcomingDates, err := dbGetUpcomingDates(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	uploads, err := dbGetUploads(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	vaccine, err := dbGetVaccine(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	profileInfo.Patient = patient
	profileInfo.AllergiesInfo = allergies
	profileInfo.ComplaintInfo = complaint
	profileInfo.DoctorNotes = doctorNotes
	profileInfo.FamilyHistory = familyHistory
	profileInfo.Medications = medications
	profileInfo.UpcomingDates = upcomingDates
	profileInfo.Uploads = uploads
	profileInfo.Vaccine = vaccine
	response.Data = profileInfo
	json.NewEncoder(w).Encode(response)
}

func editBasicInfo(w http.ResponseWriter, r *http.Request) {
	var (
		request  BasicInfoStruct
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbEditBasicInfo(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		println(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func editCurrentComplaint(w http.ResponseWriter, r *http.Request) {
	var (
		request  EditCurrentComplaintStruct
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbEditCurrentComplaint(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addNewComplaint(w http.ResponseWriter, r *http.Request) {
	var (
		request  ComplaintInfo
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddNewComplaint(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func editComplaint(w http.ResponseWriter, r *http.Request) {
	var (
		request  ComplaintInfo
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbEditComplaint(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func deleteComplaint(w http.ResponseWriter, r *http.Request) {
	var (
		request  ComplaintInfo
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteComplaint(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func uploadNewFile(w http.ResponseWriter, r *http.Request) {
	var (
		request  Uploads
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbUploadNewFile(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addUpcomingDate(w http.ResponseWriter, r *http.Request) {
	var (
		request  UpcomingDates
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddUpcomingDate(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func deleteDate(w http.ResponseWriter, r *http.Request) {
	var (
		request  UpcomingDates
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteDate(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addNewVaccine(w http.ResponseWriter, r *http.Request) {
	var (
		request  Vaccine
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddNewVaccine(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func deleteVaccine(w http.ResponseWriter, r *http.Request) {
	var (
		request  Vaccine
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteVaccine(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addNewMedication(w http.ResponseWriter, r *http.Request) {
	var (
		request  Medications
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddNewMedication(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func editMedication(w http.ResponseWriter, r *http.Request) {
	var (
		request  Medications
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbEditMedication(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func deleteMedication(w http.ResponseWriter, r *http.Request) {
	var (
		request  Medications
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteMedication(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addNewAllergy(w http.ResponseWriter, r *http.Request) {
	var (
		request  AllergiesInfo
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddNewAllergy(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func editAllergy(w http.ResponseWriter, r *http.Request) {
	var (
		request  AllergiesInfo
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbEditAllergy(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func deleteAllergy(w http.ResponseWriter, r *http.Request) {
	var (
		request  AllergiesInfo
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteAllergy(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addNewFamHistory(w http.ResponseWriter, r *http.Request) {
	var (
		request  FamilyHistory
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddNewFamHistory(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func editFamHistory(w http.ResponseWriter, r *http.Request) {
	var (
		request  FamilyHistory
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbEditFamHistory(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func deleteFamHistory(w http.ResponseWriter, r *http.Request) {
	var (
		request  FamilyHistory
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteFamHistory(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	var (
		request  Patient
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddNewUser(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	var (
		request  DownFile
		response DownFileREs
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	file, err := dbDownloadFile(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	response.File = file
	json.NewEncoder(w).Encode(response)
}

func profilePicture(w http.ResponseWriter, r *http.Request) {
	var (
		request  ProfilePic
		response ProfilePicRes
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	profile, err := dbProfilePicture(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	response.Picture = profile
	json.NewEncoder(w).Encode(response)
}

func defaultProfilePicture(w http.ResponseWriter, r *http.Request) {
	var (
		request  ProfilePic
		response ProfilePicRes
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	profile, err := dbDefaultProfilePicture(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	response.Picture = profile
	json.NewEncoder(w).Encode(response)
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	var (
		request  DownFile
		response DownFileREs
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteFile(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func addDoctorRemark(w http.ResponseWriter, r *http.Request) {
	var (
		request  DoctorNotes
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbAddDoctorRemark(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func editDoctorRemark(w http.ResponseWriter, r *http.Request) {
	var (
		request  DoctorNotes
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbEditDoctorRemark(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func deleteDoctorRemark(w http.ResponseWriter, r *http.Request) {
	var (
		request  DoctorNotes
		response StandardResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	err = dbDeleteDoctorRemark(request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	json.NewEncoder(w).Encode(response)
}

func qr(w http.ResponseWriter, r *http.Request) {
	var (
		request     StandardIdRequest
		response    StandardResponse
		profileInfo ProfileInfo
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	patient, err := dbGetPatientInfo(request.Id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Message = "Operation Succeeded"
	profileInfo.Patient = patient
	response.Data = profileInfo
	json.NewEncoder(w).Encode(response)
}
