package database

import (
	"AppointmentSummmary_Assignment/models"
	//"database/sql"
	"fmt"
	"time"
)

func ReadDataForDate(date string) ([]models.AppointmentDetails, error) {
	// layout := "2006-01-02"
	// startOfDay, err := time.Parse(layout, date)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid date format. expected YYYY-MM-DD: %w", err)
	// }
	// endOfDay := startOfDay.Add(24 * time.Hour)

	query := `
	SELECT
		a.appointment_id,
		c.center_id,
		c.center_name,
		d.doctor_id,
		d.name AS doctor_name,
		d.mobile AS doctor_mobile,
		p.patient_id,
		CONCAT(p.salutation, ' ', p.name) AS patient_name,
		a.appointment_start_dttm,
		a.appointment_end_ddttm,
		a.treatment_category
	FROM appointments a
	JOIN centers c ON a.center_id = c.center_id
	JOIN doctors d ON a.doctorstaff_id = d.doctor_id
	JOIN patients p ON a.patient_id = p.patient_id
	WHERE a.appointment_status = 'S'
  AND TO_DATE(SPLIT_PART(a.appointment_start_dttm, ' ', 1), 'MM/DD/YYYY') = $1
	ORDER BY d.doctor_id, c.center_id, a.appointment_start_dttm;
	`

	rows, err := DB.Query(query, date)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var result []models.AppointmentDetails

	for rows.Next() {
		var a models.AppointmentDetails
		var startStr, endStr string

		err := rows.Scan(
			&a.AppointmentID,
			&a.CenterID,
			&a.CenterName,
			&a.DoctorID,
			&a.DoctorName,
			&a.DoctorMobile,
			&a.PatientID,
			&a.PatientName,
			&startStr,
			&endStr,
			&a.TreatmentCategory,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		// a.AppointmentStartDttm, _ = time.Parse("2006-01-02 15:04", startStr)
		// a.AppointmentEndDttm, _ = time.Parse("2006-01-02 15:04", endStr)
		a.AppointmentStartDttm, _ = time.Parse("1/2/2006 15:04", startStr)
		a.AppointmentEndDttm, _ = time.Parse("1/2/2006 15:04", endStr)

		result = append(result, a)
	}

	return result, nil
}
