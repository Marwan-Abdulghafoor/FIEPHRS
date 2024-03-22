package main

import "net/http"

func main() {
	initializing()

	http.HandleFunc("/login", login)
	http.HandleFunc("/adminLogin", adminLogin)
	http.HandleFunc("/checkSession", checkSession)
	http.HandleFunc("/checkAdminSession", checkAdminSession)
	http.HandleFunc("/addNewUser", addNewUser)
	http.HandleFunc("/qr", qr)

	http.HandleFunc("/getProfileInfo", getProfileInfo)
	http.HandleFunc("/editBasicInfo", editBasicInfo)
	http.HandleFunc("/editCurrentComplaint", editCurrentComplaint)
	http.HandleFunc("/addNewComplaint", addNewComplaint)
	http.HandleFunc("/editComplaint", editComplaint)
	http.HandleFunc("/deleteComplaint", deleteComplaint)
	http.HandleFunc("/uploadNewFile", uploadNewFile)
	http.HandleFunc("/deleteFile", deleteFile)
	http.HandleFunc("/downloadFile", downloadFile)
	http.HandleFunc("/profilePicture", profilePicture)
	http.HandleFunc("/defaultProfilePicture", defaultProfilePicture)
	http.HandleFunc("/addUpcomingDate", addUpcomingDate)
	http.HandleFunc("/deleteDate", deleteDate)
	http.HandleFunc("/addNewVaccine", addNewVaccine)
	http.HandleFunc("/deleteVaccine", deleteVaccine)
	http.HandleFunc("/addNewMedication", addNewMedication)
	http.HandleFunc("/editMedication", editMedication)
	http.HandleFunc("/deleteMedication", deleteMedication)
	http.HandleFunc("/addNewAllergy", addNewAllergy)
	http.HandleFunc("/editAllergy", editAllergy)
	http.HandleFunc("/deleteAllergy", deleteAllergy)
	http.HandleFunc("/addNewFamHistory", addNewFamHistory)
	http.HandleFunc("/editFamHistory", editFamHistory)
	http.HandleFunc("/deleteFamHistory", deleteFamHistory)
	http.HandleFunc("/addDoctorRemark", addDoctorRemark)
	http.HandleFunc("/editDoctorRemark", editDoctorRemark)
	http.HandleFunc("/deleteDoctorRemark", deleteDoctorRemark)

	http.ListenAndServe(":5050", nil)
}
