package main

import (
	"AppointmentSummmary_Assignment/database"
	"AppointmentSummmary_Assignment/sender"
	"fmt"
	"log"
	"os"
	//"time"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found. Falling back to OS environment variables")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <YYYY-MM-DD>")
	}
	date := os.Args[1]

	database.InitDB()
	defer database.DB.Close()
	if len(os.Args) != 2 {
		log.Fatalf("Usage: go run main.go YYYY-MM-DD")
	}
	// date := os.Args[1]

	// 1. Connect to DB
	database.InitDB()

	// 2. Create schema
	if err := database.CreateSchema(); err != nil {
		log.Fatalf("❌ Schema creation failed: %v", err)
	}

	// 3. Load and insert centers, doctors, patients
	if err := database.SeedStaticData(); err != nil {
		log.Fatalf("❌ Seeding static data failed: %v", err)
	}
	fmt.Println("✅ Seeded centers, doctors, and patients")

	// 4. Load and insert appointments
	appts, err := database.LoadAppointments()
	if err != nil {
		log.Fatalf("❌ Failed to load appointments CSV: %v", err)
	}
	fmt.Printf("✅ Parsed %d appointments from CSV\n", len(appts))

	if err := database.InsertAppointments(appts); err != nil {
		log.Fatalf("❌ Failed to insert appointments: %v", err)
	}
	fmt.Println("✅ Appointments inserted into DB")

	// 5. Read appointments for given date
	apptsForDate, err := database.ReadDataForDate(date)
	if err != nil {
		log.Fatalf("❌ Failed to fetch appointment details: %v", err)
	}
	fmt.Printf("✅ Loaded %d appointments for date %s\n", len(apptsForDate), date)

	// 6. Generate and insert messages
	if len(apptsForDate) > 0 {
		if err := sender.CreateAndScheduleSummaryAppointmentMessages(apptsForDate); err != nil {
			log.Fatalf("❌ Failed to generate summary messages: %v", err)
		}
		fmt.Println("✅ Summary messages generated and inserted successfully.")
	} else {
		fmt.Println("ℹ️ No scheduled appointments found for the given date.")
	}
}
