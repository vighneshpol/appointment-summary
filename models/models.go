package models

import (
	"time"
)

type Center struct {
	ID   int64  `db:"center_id"`
	Name string `db:"center_name"`
}

type Patient struct {
	ID         int64  `db:"patient_id"`
	Salutation string `db:"salutation"`
	Name       string `db:"name"`
	Mobile     string `db:"mobile"`
}

type Doctor struct {
	ID     int64  `db:"doctor_id"`
	Name   string `db:"name"`
	Mobile string `db:"mobile"`
}

type Appointment struct {
	ID                 int64  `db:"appointment_id"`
	CenterID           int64  `db:"center_id"`
	DoctorID           int64  `db:"doctorstaff_id"`
	PatientID          int64  `db:"patient_id"`
	StartTimeRaw       string `db:"appointment_start_dttm"` // raw CSV datetime string
	EndTimeRaw         string `db:"appointment_end_ddttm"`
	Status             string `db:"appointment_status"`
	TreatmentCategory  string `db:"treatment_category"`
}

type SummaryInputData struct {
	DoctorID          int64
	DoctorName        string
	DoctorMobile      string
	PatientName       string
	PatientSalutation string
	StartTime         string
	EndTime           string
	TreatmentCategory string
	CenterID          int64
	CenterName        string
	AppointmentDate   string
}

type DoctorMessage struct {
	DoctorID     int64
	DoctorMobile string
	Message      string
}

type CenterMessage struct {
	CenterID int64
	Message  string
}

type AppointmentDetails struct {
	AppointmentID       int64
	CenterID            int64
	CenterName          string
	DoctorID            int64
	DoctorName          string
	DoctorMobile        string
	PatientID           int64
	PatientName         string
	AppointmentStartDttm time.Time
	AppointmentEndDttm   time.Time
	TreatmentCategory   string
}
