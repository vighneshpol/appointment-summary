package database

import (
	"fmt"
)

func CreateSchema() error {
	schema := []string{
		`CREATE TABLE IF NOT EXISTS centers (
			center_id   BIGINT PRIMARY KEY,
			center_name TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS patients (
			patient_id  BIGINT PRIMARY KEY,
			salutation  TEXT,
			name        TEXT NOT NULL,
			mobile      TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS doctors (
			doctor_id   BIGINT PRIMARY KEY,
			name        TEXT NOT NULL,
			mobile      TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS appointments (
			appointment_id BIGINT PRIMARY KEY,
			center_id      BIGINT REFERENCES centers(center_id),
			doctorstaff_id BIGINT REFERENCES doctors(doctor_id),
			patient_id     BIGINT REFERENCES patients(patient_id),
			appointment_start_dttm TEXT,
			appointment_end_ddttm TEXT,
			appointment_status TEXT,
			treatment_category TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS doctor_messages (
			id SERIAL PRIMARY KEY,
			doctor_id   BIGINT,
			doctor_mobile TEXT,
			message TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS center_messages (
			id SERIAL PRIMARY KEY,
			center_id BIGINT,
			message TEXT
		);`,

		// Indexes for performance
		`CREATE INDEX IF NOT EXISTS idx_appt_status_date ON appointments (appointment_status, appointment_start_dttm);`,
		`CREATE INDEX IF NOT EXISTS idx_appt_doctor ON appointments (doctorstaff_id);`,
		`CREATE INDEX IF NOT EXISTS idx_appt_center ON appointments (center_id);`,
	}

	tx := DB.MustBegin()
	for _, stmt := range schema {
		if _, err := tx.Exec(stmt); err != nil {
			return fmt.Errorf("schema creation failed: %v", err)
		}
	}
	return tx.Commit()
}
