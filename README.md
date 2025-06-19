# 🩺 Appointment Summary Generator

This Golang-based project loads healthcare appointment data from CSV files, seeds it into a PostgreSQL database, and generates summary messages for each doctor and center for a specific date.

---

## 📁 Folder Structure

```
AppointmentSummmary_Assignment/
├── data/
│   ├── Appointment.csv
│   ├── Center.csv
│   ├── DoctorStaff.csv
│   └── Patient.csv
├── database/
│   ├── connect.go
│   ├── loader.go
│   ├── schema.go
│   └── seeder.go
├── sender/
│   └── sender.go
├── models/
│   └── models.go
├── config/
│       └── config.go
├── go.mod
└── main.go
```

---

##  Features

- Parses appointment, doctor, center, and patient data from CSV files
- Seeds data into PostgreSQL with FK constraints
- Dynamically loads summary messages for a given date
- Generates:
  - Doctor-wise appointment summaries
  - Center-wise appointment summaries

---

## ⚙️ Prerequisites

- Go 1.18+
- PostgreSQL installed and running
- [Optional] VS Code for development

---

## 🛠️ Environment Variables

Create a `.env` file in the root (optional if using fallback defaults):

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=appointment_summary
DB_SSLMODE=disable
```

---

## 📦 Setup Instructions

### 1️⃣ Clone the repo

```bash
git clone https://github.com/vighneshpol/appointment-summary.git
cd AppointmentSummmary_Assignment
```

### 2️⃣ Get dependencies

```bash
go mod tidy
```

### 3️⃣ Create PostgreSQL DB

```sql
CREATE DATABASE appointment_summary;
```

Or from terminal:
```bash
psql -U postgres -c "CREATE DATABASE appointment_summary;"
```

### 4️⃣ Add your CSV files in `/data/` folder

- Ensure all 4 files are present: `Appointment.csv`, `DoctorStaff.csv`, `Center.csv`, `Patient.csv`
- Format must match expected columns (consult `models.go`)

---

## 🚀 Running the Project

```bash
go run main.go 2025-05-12
```

**Replace `2025-05-12`** with any desired date in format `YYYY-MM-DD`

---

## ✅ Expected Output

```bash
✅ Connected to PostgreSQL database successfully.
✅ Parsed 500000 appointments from CSV
⏳ Inserting centers...
✅ Inserted centers...
⏳ Inserting doctors...
✅ Inserted doctors...
⏳ Inserting patient...
✅ Inserted patient...
⏳ Inserting appointment...
✅ Inserting appointments done...
✅ Seeded centers, doctors, and patients
✅ Loaded 95625 appointments for date 2025-05-12
✅ Summary messages generated and inserted successfully.
```

---

## 🧪 How to Test in VS Code Terminal

1. Open VS Code terminal:  
   `Terminal → New Terminal` or `Ctrl + \``

2. Run the project:

```bash
go run main.go 2025-05-12
```

3. Output should show ✅ messages above

4. To verify messages:

```sql
SELECT * FROM doctor_messages LIMIT 5;
SELECT * FROM center_messages LIMIT 5;
```

---

## 💡 Troubleshooting

| Issue | Fix |
|-------|-----|
| `invalid date format` | Use `YYYY-MM-DD` format like `2025-05-12` |
| `foreign key constraint error` | Seed doctors/centers/patients before appointments |
| `no messages generated` | Check appointment CSV: dates, status='S', time format |
| `psql error` | Ensure PostgreSQL is installed and running |

---

## 🔐 Security Note

This project supports environment variables via `.env` for sensitive info like DB password.  
Avoid committing `.env` to source control.

---

## 👨‍💻 Author

Built by Vighnesh Mukund Pol

---

## 📝 License

MIT License. Use freely for educational and internal purposes.

