package main

import (
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
)

func dbGetOldSessions(username string) []string {
	var sessions []string
	query := `SELECT sessionid FROM sessions 
	WHERE username = ? AND NOW() > time_left`
	rows, err := db.Query(query, username)
	if err != nil {
		WriteLog("dbGetSessions :", err.Error())
		return sessions
	}
	defer rows.Close()
	for rows.Next() {
		var sessionId string
		rows.Scan(&sessionId)
		sessions = append(sessions, sessionId)
	}
	return sessions
}

func dbDeleteSession(sessionId string) {
	query := `DELETE FROM sessions WHERE sessionid = ?`
	_, err := db.Exec(query, sessionId)
	if err != nil {
		WriteLog("dbDeleteSession *DELETE SESSIONID* -> :", err.Error())
	}
}

func checkValidSessions(username string) {
	sessions := dbGetOldSessions(username)
	for _, val := range sessions {
		dbDeleteSession(val)
	}
}

func generateSessionAndTimeLeft(username string) (string, string) {
	sessionByteA, _ := uuid.NewRandom()
	sessionStringA := sessionByteA.String()
	sessionByteB, _ := uuid.NewRandom()
	sessionStringB := sessionByteB.String()
	sessionString := sessionStringA + sessionStringB
	timeLeft := time.Now().Add(time.Hour * 8).Format(time.RFC1123)
	session := strings.ReplaceAll(sessionString, "-", "")
	return session, timeLeft
}

func dbInsertSession(username, session, timeLeft string) error {
	query := `INSERT INTO sessions (username, sessionid, time_left) 
	VALUES (?,?,DATE_ADD(NOW(), INTERVAL 8 HOUR))`
	_, err := db.Exec(query, username, session)
	if err != nil {
		WriteLog("dbInsertSession :", err.Error())
		return err
	}
	return nil
}

func dbCheckSession(sessionid string) (string, error) {
	var username string
	query := `SELECT username FROM sessions 
	WHERE sessionid = ? AND NOW() < time_left`
	err := db.QueryRow(query, sessionid).Scan(&username)
	if err == sql.ErrNoRows {
		return username, err
	}
	if err != nil {
		WriteLog("dbCheckSession *DELETE SESSIONID* -> :", err.Error())
		return username, err
	}
	return username, nil
}
