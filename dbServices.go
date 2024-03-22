package main

import "database/sql"

func checkUser(username, password string) bool {
	var user, pass string
	err := db.QueryRow(`SELECT username, password FROM patients
	 WHERE username = ? and password = md5(?);`, username, password).Scan(&user, &pass)
	if err != nil {
		if err != sql.ErrNoRows {
			WriteLog(err.Error())
		}
		return false
	}

	return true
}

func checkAdmin(username, password string) bool {
	var user, pass string
	err := db.QueryRow(`SELECT admin, password FROM admins
	 WHERE admin = ? and password = md5(?);`, username, password).Scan(&user, &pass)
	if err != nil {
		if err != sql.ErrNoRows {
			WriteLog(err.Error())
		}
		return false
	}

	return true
}

func dbGetUserId(username string) (int, error) {
	var id int
	query := "SELECT id FROM patients WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&id)
	if err != nil {
		WriteLog("dbGetUserId : ", err.Error())
		return id, err
	}
	return id, err
}

func dbGetAdminId(username string) (int, error) {
	var id int
	query := "SELECT id FROM admins WHERE admin = ?"
	err := db.QueryRow(query, username).Scan(&id)
	if err != nil {
		WriteLog("dbGetAdminId : ", err.Error())
		return id, err
	}
	return id, err
}

func dbGetPatientInfo(id int) (Patient, error) {
	var patient Patient
	query := `SELECT * FROM patients WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&patient.Id,
		&patient.Username, &patient.Password, &patient.Email, &patient.Name,
		&patient.Sex, &patient.Address, &patient.DOB,
		&patient.BType, &patient.Allergies, &patient.Reactions,
		&patient.CMC, &patient.FamilyCMC, &patient.MentalCondition,
		&patient.Surgeries, &patient.Complaint, &patient.Marital,
		&patient.Smoke, &patient.Contacts, &patient.Number,
		&patient.Height, &patient.Weight, &patient.IceInstructions,
		&patient.Lang, &patient.Role, &patient.Class, &patient.Profile,
		&patient.ProfileLink)
	if err != nil {
		println(id)
		WriteLog("dbGetPatientInfo : ", err.Error())
		return patient, err
	}
	return patient, err
}

func dbAddNewUser(request Patient) error {
	query := `INSERT INTO patients (name, username, password,
		dob, btype, address, sex, marital, allergies, height, 
		weight, smoke, surgeries, contacts, cmc, mentalconditions,
		ice_instructions, complaint, lang, class, role, number, 
		email, familycmc, reactions, profile, profile_link) VALUES (?, ?, md5(?), 
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Username,
		request.Password, request.DOB, request.BType, request.Address,
		request.Sex, request.Marital, request.Allergies, request.Height,
		request.Weight, request.Smoke, request.Surgeries, request.Contacts,
		request.CMC, request.MentalCondition, request.IceInstructions,
		request.Complaint, request.Lang, request.Class, request.Role,
		request.Number, request.Email, request.FamilyCMC, request.Reactions,
		request.Profile, request.ProfileLink)
	if err != nil {
		WriteLog("dbAddNewUser : ", err.Error())
		return err
	}
	return err
}

func dbEditBasicInfo(request BasicInfoStruct) error {
	query := `UPDATE patients SET name = ?, dob = ?,
	btype = ?, address = ?, sex = ?, marital = ?,
	allergies = ?, height = ?, weight = ?, smoke = ?,
	surgeries = ?, contacts = ?, cmc = ?, mentalconditions = ?,
	ice_instructions = ?, profile = ?, profile_link = ? WHERE id = ?`
	_, err := db.Exec(query, request.Name, request.DOB,
		request.Btype, request.Address, request.Gender, request.MartialStatus,
		request.Allergies, request.Height, request.Weight, request.Smoke,
		request.Surgeries, request.Contacts, request.CMC,
		request.MentalCon, request.IceInstructions, request.Profile, request.ProfileLink, request.Id)
	if err != nil {
		WriteLog("dbEditBasicInfo : ", err.Error())
		return err
	}
	return err
}

func dbEditCurrentComplaint(request EditCurrentComplaintStruct) error {
	query := `UPDATE patients SET complaint = ? WHERE id = ?`
	_, err := db.Exec(query, request.Complaint, request.Id)
	if err != nil {
		WriteLog("dbEditCurrentComplaint : ", err.Error())
		return err
	}
	return err
}

func dbGetAllergies(id int) ([]AllergiesInfo, error) {
	var allergiesInfos []AllergiesInfo
	query := `SELECT * FROM allergies WHERE patient_id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetAllergies : ", err.Error())
		return allergiesInfos, err
	}
	for rows.Next() {
		var allergiesInfo AllergiesInfo
		rows.Scan(&allergiesInfo.Id, &allergiesInfo.Name,
			&allergiesInfo.Reaction, &allergiesInfo.PatientId)
		allergiesInfos = append(allergiesInfos, allergiesInfo)
	}

	return allergiesInfos, err
}

func dbAddNewAllergy(request AllergiesInfo) error {
	query := `INSERT INTO allergies (name, reaction, patient_id)
	VALUES (?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Reaction, request.PatientId)
	if err != nil {
		WriteLog("dbAddNewAllergy : ", err.Error())
		return err
	}
	return err
}

func dbEditAllergy(request AllergiesInfo) error {
	query := `UPDATE allergies SET name = ?, 
	reaction = ? WHERE id = ?`
	_, err := db.Exec(query, request.Name, request.Reaction, request.Id)
	if err != nil {
		WriteLog("dbEditAllergy : ", err.Error())
		return err
	}
	return err
}

func dbDeleteAllergy(request AllergiesInfo) error {
	query := `DELETE FROM allergies WHERE id = ?`
	_, err := db.Exec(query, request.Id)
	if err != nil {
		WriteLog("dbDeleteAllergy : ", err.Error())
		return err
	}
	return err
}

func dbGetComplaint(id int) ([]ComplaintInfo, error) {
	var complaintInfos []ComplaintInfo
	query := `SELECT * FROM complaint WHERE patient_id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetComplaint : ", err.Error())
		return complaintInfos, err
	}
	for rows.Next() {
		var complaintInfo ComplaintInfo
		rows.Scan(&complaintInfo.Id, &complaintInfo.Name,
			&complaintInfo.Descriprion, &complaintInfo.Date, &complaintInfo.PatientId)
		complaintInfos = append(complaintInfos, complaintInfo)
	}
	return complaintInfos, err
}
func dbAddNewComplaint(request ComplaintInfo) error {
	query := `INSERT INTO complaint (name, description, date, patient_id) 
	VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Descriprion,
		request.Date, request.PatientId)
	if err != nil {
		WriteLog("dbAddNewComplaint : ", err.Error())
		return err
	}
	return err
}

func dbEditComplaint(request ComplaintInfo) error {
	query := `UPDATE complaint SET name = ?, date = ?,
	description = ? WHERE id = ?`
	_, err := db.Exec(query, request.Name, request.Date,
		request.Descriprion, request.Id)
	if err != nil {
		WriteLog("dbEditComplaint : ", err.Error())
		return err
	}
	return err
}

func dbDeleteComplaint(request ComplaintInfo) error {
	query := `DELETE FROM complaint WHERE id = ?`
	_, err := db.Exec(query, request.Id)
	if err != nil {
		WriteLog("dbDeleteComplaint : ", err.Error())
		return err
	}
	return err
}

func dbGetDoctorNotes(id int) ([]DoctorNotes, error) {
	var doctorNotes []DoctorNotes
	query := `SELECT * FROM doctor_notes WHERE patient_id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetDoctorNotes : ", err.Error())
		return doctorNotes, err
	}
	for rows.Next() {
		var doctorNote DoctorNotes
		rows.Scan(&doctorNote.Id, &doctorNote.Name,
			&doctorNote.Number, &doctorNote.Location,
			&doctorNote.Notes, &doctorNote.Speciality,
			&doctorNote.PatientId)
		doctorNotes = append(doctorNotes, doctorNote)
	}
	return doctorNotes, err
}

func dbAddDoctorRemark(request DoctorNotes) error {
	query := `INSERT INTO doctor_notes (name, number, location,
		notes, Speciality, patient_id) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Number, request.Location,
		request.Notes, request.Speciality, request.PatientId)
	if err != nil {
		WriteLog("dbAddDoctorRemark : ", err.Error())
		return err
	}
	return err
}

func dbEditDoctorRemark(request DoctorNotes) error {
	query := `UPDATE doctor_notes SET name = ?,
	 number = ?, location = ?, notes = ?,
	 Speciality = ? WHERE id = ?`
	_, err := db.Exec(query, request.Name, request.Number, request.Location,
		request.Notes, request.Speciality, request.Id)
	if err != nil {
		WriteLog("dbEditDoctorRemark : ", err.Error())
		return err
	}
	return err
}

func dbDeleteDoctorRemark(request DoctorNotes) error {
	query := `DELETE FROM doctor_notes WHERE id = ?`
	_, err := db.Exec(query, request.Id)
	if err != nil {
		WriteLog("dbDeleteDoctorRemark : ", err.Error())
		return err
	}
	return err
}

func dbGetFamilyHistory(id int) ([]FamilyHistory, error) {
	var familyHistories []FamilyHistory
	query := `SELECT * FROM familyhistory WHERE patient = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetFamilyHistory", err.Error())
		return familyHistories, err
	}
	for rows.Next() {
		var familyHistory FamilyHistory
		rows.Scan(&familyHistory.Id, &familyHistory.Name,
			&familyHistory.Relationship, &familyHistory.MedicalConditions,
			&familyHistory.PatientId)
		familyHistories = append(familyHistories, familyHistory)
	}
	return familyHistories, err
}

func dbAddNewFamHistory(request FamilyHistory) error {
	query := `INSERT INTO familyhistory (Name, Relationship,
		Medicalconditions, patient) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Relationship,
		request.MedicalConditions, request.PatientId)
	if err != nil {
		WriteLog("dbAddNewFamHistory", err.Error())
		return err
	}
	return err
}

func dbEditFamHistory(request FamilyHistory) error {
	query := `UPDATE familyhistory SET Name = ?, 
	Relationship = ?, Medicalconditions = ? WHERE id = ?`
	_, err := db.Exec(query, request.Name, request.Relationship,
		request.MedicalConditions, request.Id)
	if err != nil {
		WriteLog("dbEditFamHistory", err.Error())
		return err
	}
	return err
}

func dbDeleteFamHistory(request FamilyHistory) error {
	query := `DELETE FROM familyhistory WHERE id = ?`
	_, err := db.Exec(query, request.Id)
	if err != nil {
		WriteLog("dbDeleteFamHistory", err.Error())
		return err
	}
	return err
}

func dbGetMedications(id int) ([]Medications, error) {
	var medications []Medications
	query := `SELECT * FROM medications WHERE patient_id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetMedications : ", err.Error())
		return medications, err
	}
	for rows.Next() {
		var medication Medications
		rows.Scan(&medication.Id, &medication.Name,
			&medication.Dosage, &medication.Frequency,
			&medication.PatientId)
		medications = append(medications, medication)
	}
	return medications, err
}

func dbAddNewMedication(request Medications) error {
	query := `INSERT INTO medications (name,
		dosage, frequency, patient_id) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Dosage,
		request.Frequency, request.PatientId)
	if err != nil {
		WriteLog("dbAddNewMedication : ", err.Error())
		return err
	}
	return err
}

func dbEditMedication(request Medications) error {
	query := `UPDATE medications SET name = ?,
		dosage = ?, frequency = ? WHERE id = ?`
	_, err := db.Exec(query, request.Name, request.Dosage,
		request.Frequency, request.Id)
	if err != nil {
		WriteLog("dbEditMedication : ", err.Error())
		return err
	}
	return err
}

func dbDeleteMedication(request Medications) error {
	query := `DELETE FROM medications WHERE id = ?`
	_, err := db.Exec(query, request.Id)
	if err != nil {
		WriteLog("dbDeleteMedication : ", err.Error())
		return err
	}
	return err
}

func dbGetUpcomingDates(id int) ([]UpcomingDates, error) {
	var upcomingDates []UpcomingDates
	query := `SELECT * FROM upcoming_dates WHERE patient_id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetUpcomingDates : ", err.Error())
		return upcomingDates, err
	}
	for rows.Next() {
		var upcomingDate UpcomingDates
		rows.Scan(&upcomingDate.Id, &upcomingDate.Name,
			&upcomingDate.Date, &upcomingDate.Time,
			&upcomingDate.PatientId)
		upcomingDates = append(upcomingDates, upcomingDate)
	}
	return upcomingDates, err
}

func dbAddUpcomingDate(request UpcomingDates) error {
	query := `INSERT INTO upcoming_dates (name, date, 
		time, patient_id) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Date,
		request.Time, request.PatientId)
	if err != nil {
		WriteLog("dbAddUpcomingDate : ", err.Error())
		return err
	}
	return err
}

func dbDeleteDate(request UpcomingDates) error {
	query := `DELETE FROM upcoming_dates WHERE id = ?`
	_, err := db.Exec(query, request.Id)
	if err != nil {
		WriteLog("dbDeleteDate : ", err.Error())
		return err
	}
	return err
}

func dbGetUploads(id int) ([]Uploads, error) {
	var uploads []Uploads
	query := `SELECT * FROM uploads WHERE patient_id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetUploads : ", err.Error())
		return uploads, err
	}
	for rows.Next() {
		var upload Uploads
		rows.Scan(&upload.Id, &upload.Filename,
			&upload.Name, &upload.Size,
			&upload.PatientId, &upload.File)
		uploads = append(uploads, upload)
	}
	return uploads, err
}

func dbUploadNewFile(request Uploads) error {
	query := `INSERT INTO uploads (filename, name,
		 size, patient_id, file) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, request.Filename, request.Name,
		request.Size, request.PatientId, request.File)
	if err != nil {
		WriteLog("dbUploadNewFile : ", err.Error())
		return err
	}
	return err
}

func dbGetVaccine(id int) ([]Vaccine, error) {
	var vaccines []Vaccine
	query := `SELECT * FROM vaccine WHERE patient_id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		WriteLog("dbGetVaccine : ", err.Error())
		return vaccines, err
	}
	for rows.Next() {
		var vaccine Vaccine
		rows.Scan(&vaccine.Id, &vaccine.Name,
			&vaccine.Details, &vaccine.Date, &vaccine.PatientId)
		vaccines = append(vaccines, vaccine)
	}
	return vaccines, err
}

func dbAddNewVaccine(request Vaccine) error {
	query := `INSERT INTO vaccine (name,
		details, date, patient_id) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, request.Name, request.Details,
		request.Date, request.PatientId)
	if err != nil {
		WriteLog("dbAddNewVaccine : ", err.Error())
		return err
	}
	return err
}

func dbDeleteVaccine(request Vaccine) error {
	query := `DELETE FROM vaccine WHERE id = ?`
	_, err := db.Exec(query, request.Id)
	if err != nil {
		WriteLog("dbDeleteVaccine : ", err.Error())
		return err
	}
	return err
}

func dbDownloadFile(request DownFile) ([]byte, error) {
	var file []byte
	query := `SELECT file FROM uploads WHERE id = ?`
	err := db.QueryRow(query, request.FileID).Scan(&file)
	if err != nil {
		WriteLog("dbDownloadFile : ", err.Error())
		return file, err
	}
	return file, err
}

func dbProfilePicture(request ProfilePic) ([]byte, error) {
	var profile []byte
	query := `SELECT profile FROM patients WHERE id = ?`
	err := db.QueryRow(query, request.PicID).Scan(&profile)
	if err != nil {
		WriteLog("dbProfilePicture : ", err.Error())
		return profile, err
	}
	return profile, err
}

func dbDefaultProfilePicture(request ProfilePic) ([]byte, error) {
	var profile []byte
	query := `SELECT profile FROM default_profile WHERE id = ?`
	err := db.QueryRow(query, request.PicID).Scan(&profile)
	if err != nil {
		WriteLog("dbDefaultProfilePicture : ", err.Error())
		return profile, err
	}
	return profile, err
}

func dbDeleteFile(request DownFile) error {
	query := `DELETE FROM uploads WHERE id = ?`
	_, err := db.Exec(query, request.FileID)
	if err != nil {
		WriteLog("dbDeleteFile : ", err.Error())
		return err
	}
	return err
}
