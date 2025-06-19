package sender

import (
	"AppointmentSummmary_Assignment/database"
	"AppointmentSummmary_Assignment/models"
	"fmt"
	"sort"
	"strings"
	"time"
)

func CreateAndScheduleSummaryAppointmentMessages(appointments []models.AppointmentDetails) error {
	db := database.DB

	// Group by (DoctorID, CenterID)
	type docCenterKey struct {
		DoctorID int64
		CenterID int64
	}
	doctorCenterMap := make(map[docCenterKey][]models.AppointmentDetails)

	// Also group by CenterID for center summary
	centerMap := make(map[int64]map[int64]int) // CenterID → map[DoctorID]appointmentCount
	centerTotal := make(map[int64]int)

	for _, a := range appointments {
		key := docCenterKey{DoctorID: a.DoctorID, CenterID: a.CenterID}
		doctorCenterMap[key] = append(doctorCenterMap[key], a)

		if _, exists := centerMap[a.CenterID]; !exists {
			centerMap[a.CenterID] = make(map[int64]int)
		}
		centerMap[a.CenterID][a.DoctorID]++
		centerTotal[a.CenterID]++
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Format for messages
	for _ , appts := range doctorCenterMap {
		sort.Slice(appts, func(i, j int) bool {
			return appts[i].AppointmentStartDttm.Before(appts[j].AppointmentStartDttm)
		})

		date := appts[0].AppointmentStartDttm.Format("2nd Jan, 2006")
		header := fmt.Sprintf("Dr. %s's appointments on <%s> at %s: %d",
			appts[0].DoctorName, date, appts[0].CenterName, len(appts))

		var lines []string
		for _, appt := range appts {
			start := appt.AppointmentStartDttm.Format("3:04 pm")
			duration := appt.AppointmentEndDttm.Sub(appt.AppointmentStartDttm)
			durStr := formatDuration(duration)

			category := ""
			if strings.ToLower(appt.TreatmentCategory) != "not specified" {
				category = fmt.Sprintf(" (%s)", appt.TreatmentCategory)
			}

			lines = append(lines, fmt.Sprintf("%s, %s: %s%s",
				start, durStr, appt.PatientName, category))
		}

		fullMsg := header + "\n" + strings.Join(lines, "\n")
		_, err := tx.Exec(`INSERT INTO doctor_messages (doctor_id, doctor_mobile, message) VALUES ($1, $2, $3)`,
			appts[0].DoctorID, appts[0].DoctorMobile, fullMsg)
		if err != nil {
			return fmt.Errorf("inserting doctor message failed: %w", err)
		}
	}

	// Center Summary
	for centerID, docMap := range centerMap {
		sampleAppt := findSampleAppointment(appointments, centerID)
		if sampleAppt == nil {
			continue
		}
		date := sampleAppt.AppointmentStartDttm.Format("2nd Jan, 2006")
		header := fmt.Sprintf("Summary of appointments at %s on %s: %d",
			sampleAppt.CenterName, date, centerTotal[centerID])

		var lines []string
		for docID, count := range docMap {
			docName := findDoctorName(appointments, docID)
			lines = append(lines, fmt.Sprintf("Dr. %s: %d", docName, count))
		}

		fullMsg := header + "\n" + strings.Join(lines, "\n")
		_, err := tx.Exec(`INSERT INTO center_messages (center_id, message) VALUES ($1, $2)`,
			centerID, fullMsg)
		if err != nil {
			return fmt.Errorf("inserting center message failed: %w", err)
		}
	}

	return tx.Commit()
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	h := minutes / 60
	m := minutes % 60
	if h > 0 {
		return fmt.Sprintf("%dh %dm", h, m)
	}
	return fmt.Sprintf("%dm", m)
}

func findSampleAppointment(list []models.AppointmentDetails, centerID int64) *models.AppointmentDetails {
	for _, a := range list {
		if a.CenterID == centerID {
			return &a
		}
	}
	return nil
}

func findDoctorName(list []models.AppointmentDetails, doctorID int64) string {
	for _, a := range list {
		if a.DoctorID == doctorID {
			return a.DoctorName
		}
	}
	return "Unknown"
}
