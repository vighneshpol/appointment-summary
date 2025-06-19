package database

import (
	//"AppointmentSummmary_Assignment/models"
	//"database/sql"
	"fmt"
)

func SeedStaticData() error {
	centers, err := LoadCenters()
	if err != nil {
		return fmt.Errorf("failed to load centers: %w", err)
	}
	doctors, err := LoadDoctors()
	if err != nil {
		return fmt.Errorf("failed to load doctors: %w", err)
	}
	patients, err := LoadPatients()
	if err != nil {
		return fmt.Errorf("failed to load patients: %w", err)
	}
	appointments, err := LoadAppointments()
	if err != nil {
		return fmt.Errorf("failed to load appointments: %w", err)
	}

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	fmt.Println("⏳ Inserting centers...")
	// Insert Centers
	for _, c := range centers {
	_, err := tx.Exec(`INSERT INTO centers (center_id, center_name) VALUES ($1, $2) ON CONFLICT DO NOTHING`, c.ID, c.Name)
		if err != nil {
			return fmt.Errorf("insert center: %w", err)
		}
	}
	fmt.Println("✅ Inserted centers...")

	fmt.Println("⏳ Inserting doctors...")
	// Insert Doctors
	for _, d := range doctors {
		_, err := tx.Exec(`INSERT INTO doctors (doctor_id, name, mobile) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`, d.ID, d.Name, d.Mobile)
		if err != nil {
			return fmt.Errorf("insert doctor: %w", err)
		}
	}
	fmt.Println("✅ Inserted doctors...")

	fmt.Println("⏳ Inserting patient...")
	// Insert Patients
	for _, p := range patients {
		_, err := tx.Exec(`INSERT INTO patients (patient_id, salutation, name, mobile) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`, p.ID, p.Salutation, p.Name, p.Mobile)

		if err != nil {
			return fmt.Errorf("insert patient: %w", err)
		}
	}
	fmt.Println("✅ Inserted patient...")

	fmt.Println("⏳ Inserting appointment...")
	// Insert Appointments
	for _, a := range appointments {
		_, err := tx.Exec(`INSERT INTO appointments (appointment_id, center_id, doctorstaff_id, patient_id, appointment_start_dttm, appointment_end_ddttm, appointment_status, treatment_category)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING`,
			a.ID, a.CenterID, a.DoctorID, a.PatientID, a.StartTimeRaw, a.EndTimeRaw, a.Status, a.TreatmentCategory)
		if err != nil {
			return fmt.Errorf("insert appointment: %w", err)
		}
	}
	fmt.Println("✅ Inserting apointments done...")

	return tx.Commit()
}
