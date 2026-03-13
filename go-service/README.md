# Student Report PDF Microservice

A standalone Go microservice that generates downloadable PDF reports for students by consuming the existing Node.js backend API.

## What This Service Does

This service exposes a single endpoint:

```
GET /api/v1/students/:id/report
```

When called, it:

1. Fetches student data from the Node.js backend via an internal API (`/api/v1/internal/students/:id`)
2. Generates a formatted PDF report containing the student's personal, academic, guardian, and address information
3. Returns the PDF as a downloadable file

The Go service does **not** connect to the database directly. It relies entirely on the Node.js backend for data.

## Project Structure

```
go-service/
├── main.go              # Entry point, server setup, and routing
├── handlers/
│   └── report.go        # HTTP handler for the report endpoint
├── models/
│   └── student.go       # Student data structure
├── pdf/
│   └── generator.go     # PDF generation logic
├── services/
│   └── client.go        # HTTP client for the Node.js backend API
├── .env                 # Environment configuration
├── go.mod               # Go module definition
└── go.sum               # Dependency checksums
```

## Prerequisites

- **Go** (1.22 or later)
- **PostgreSQL** database running with the `school_mgmt` database set up
- **Node.js backend** running on port 5003 (or whichever port is configured)

## Configuration

Create a `.env` file in the `go-service/` directory (one is already provided):

```env
PORT=8080
BACKEND_API_URL=http://localhost:5003
SERVICE_API_KEY=<must match the SERVICE_API_KEY in the backend .env>
```

| Variable          | Description                              | Default                  |
|-------------------|------------------------------------------|--------------------------|
| `PORT`            | Port the Go service listens on           | `8080`                   |
| `BACKEND_API_URL` | URL of the Node.js backend               | `http://localhost:5003`  |
| `SERVICE_API_KEY` | Shared key for internal API auth         | *(required)*             |

## How to Run

1. Make sure PostgreSQL is running and the database is set up.

2. Start the Node.js backend:

   ```bash
   cd backend
   npm install
   npm start
   ```

3. In a separate terminal, start the Go service:

   ```bash
   cd go-service
   go mod tidy
   go run .
   ```

4. Download a student report:

   ```bash
   curl -o report.pdf http://localhost:8080/api/v1/students/1/report
   ```

   Or open `http://localhost:8080/api/v1/students/1/report` in a browser to download the PDF directly.

## Backend Changes Made

To support this microservice, the following changes were made to the Node.js backend:

- **Completed student controller** (`students-controller.js`) - Implemented all 5 CRUD handler stubs (get all, get detail, add, update, set status)
- **Added internal API route** (`/api/v1/internal/students/:id`) - A service-to-service endpoint protected by an API key middleware instead of JWT/CSRF auth
- **Added `validateServiceKey` middleware** - Checks the `X-Service-Key` header against the `SERVICE_API_KEY` environment variable
