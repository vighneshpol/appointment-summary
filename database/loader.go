package database

import (
	"AppointmentSummmary_Assignment/models"
	"encoding/csv"
	"os"
	"strconv"
	"fmt"
	"log"
	)

func LoadCenters() ([]models.Center, error) {
	f, err := os.Open(`./data/Center.csv`)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, _ := r.ReadAll()

	var centers []models.Center
	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.ParseInt(row[0], 10, 64)
		centers = append(centers, models.Center{
			ID:   id,
			Name: row[1],
		})
	}
	return centers, nil
}

func LoadDoctors() ([]models.Doctor, error) {
	f, err := os.Open(`./data/DoctorStaff.csv`)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, _ := r.ReadAll()

	var doctors []models.Doctor
	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.ParseInt(row[0], 10, 64)
		doctors = append(doctors, models.Doctor{
			ID:     id,
			Name:   row[1],
			Mobile: row[2],
		})
	}
	return doctors, nil
}

func LoadPatients() ([]models.Patient, error) {
	f, err := os.Open(`./data/Patient.csv`)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, _ := r.ReadAll()

	var patients []models.Patient
	for i, row := range records {
		if i == 0 {
			continue
		}
		id, _ := strconv.ParseInt(row[0], 10, 64)
		patients = append(patients, models.Patient{
			ID:         id,
			Salutation: row[1],
			Name:       row[2],
			Mobile:     row[3],
		})
	}
	return patients, nil
}

// func LoadAppointments() ([]models.Appointment, error) {
// 	f, err := os.Open(`./data/Appointment.csv`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()

// 	r := csv.NewReader(f)
// 	r.Read() // skip header
// 	records, _ := r.ReadAll()

// 	var appts []models.Appointment
// 	for i, row := range records {
// 		if i == 0 {
// 			continue
// 		}
// 		id, _ := strconv.ParseInt(row[0], 10, 64)
// 		centerID, _ := strconv.ParseInt(row[1], 10, 64)
// 		doctorID, _ := strconv.ParseInt(row[2], 10, 64)
// 		patientID, _ := strconv.ParseInt(row[3], 10, 64)

// 		appts = append(appts, models.Appointment{
// 			ID:                id,
// 			CenterID:          centerID,
// 			DoctorID:          doctorID,
// 			PatientID:         patientID,
// 			StartTimeRaw:      row[4],
// 			EndTimeRaw:        row[5],
// 			Status:            row[6],
// 			TreatmentCategory: row[7],
// 		})
// 	}
// 	return appts, nil
// }

func LoadAppointments() ([]models.Appointment, error) {
	f, err := os.Open("./data/Appointment.csv") // Update path if needed
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	// ✅ Skip CSV header line
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	var appts []models.Appointment
	count := 0

	for {
		row, err := r.Read()
		if err != nil {
			break // io.EOF will be caught here
		}

		id, err1 := strconv.ParseInt(row[0], 10, 64)
		centerID, err2 := strconv.ParseInt(row[1], 10, 64)
		doctorID, err3 := strconv.ParseInt(row[2], 10, 64)
		patientID, err4 := strconv.ParseInt(row[3], 10, 64)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Printf("Skipping row with bad IDs: %+v", row)
			continue
		}

		appts = append(appts, models.Appointment{
			ID:                id,
			CenterID:          centerID,
			DoctorID:          doctorID,
			PatientID:         patientID,
			StartTimeRaw:      row[4],
			EndTimeRaw:        row[5],
			Status:            row[6],
			TreatmentCategory: row[7],
		})
		count++
	}

	log.Printf("✅ Parsed %d appointments from CSV", count)
	return appts, nil
}
func InsertAppointments(appts []models.Appointment) error {
	tx := DB.MustBegin()

	for _, a := range appts {
		_, err := tx.Exec(`
			INSERT INTO appointments (
				appointment_id,
				center_id,
				doctorstaff_id,
				patient_id,
				appointment_start_dttm,
				appointment_end_ddttm,
				appointment_status,
				treatment_category
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (appointment_id) DO NOTHING
		`, a.ID, a.CenterID, a.DoctorID, a.PatientID, a.StartTimeRaw, a.EndTimeRaw, a.Status, a.TreatmentCategory)

		if err != nil {
			return fmt.Errorf("failed to insert appointment %d: %w", a.ID, err)
		}
	}

	return tx.Commit()
}
