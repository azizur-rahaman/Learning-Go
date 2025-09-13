package models

import (
	"azizur/rest-api/db"
)

type Registration struct {
	ID      int64 `json:"id"`
	EventID int64 `json:"eventId" binding:"required"`
	UserID  int64 `json:"userId"`
}

func (r Registration) Save() error {
	query := `
	INSERT INTO registrations(event_id, user_id)
	VALUES(?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(r.EventID, r.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	r.ID = id

	return err
}

func (r Registration) Delete() error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.EventID, r.UserID)
	return err
}

func GetRegistrationsByEvent(eventID int64) ([]Registration, error) {
	query := `SELECT id, event_id, user_id FROM registrations WHERE event_id = ?`

	rows, err := db.DB.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registrations []Registration
	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.EventID, &registration.UserID)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	return registrations, nil
}

func GetRegistrationsByUser(userID int64) ([]Registration, error) {
	query := `SELECT id, event_id, user_id FROM registrations WHERE user_id = ?`

	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registrations []Registration
	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.EventID, &registration.UserID)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	return registrations, nil
}

func CheckRegistrationExists(eventID, userID int64) (bool, error) {
	query := `SELECT COUNT(*) FROM registrations WHERE event_id = ? AND user_id = ?`

	var count int
	err := db.DB.QueryRow(query, eventID, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetEventWithRegistrations(eventID int64) (*Event, []Registration, error) {
	event, err := GetEventById(eventID)
	if err != nil {
		return nil, nil, err
	}

	registrations, err := GetRegistrationsByEvent(eventID)
	if err != nil {
		return nil, nil, err
	}

	return event, registrations, nil
}
