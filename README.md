# Student Grade Management System API

A backend-only REST API built in **Go (Golang)** for managing student grades, courses, and GPA calculation, similar to university portals such as Canvas or Blackboard.

This is an **individual project**, implemented strictly under the given technical and functional constraints.

## Project Overview

The system supports **exactly three roles**:

- Admin
- Teacher
- Student

Role enforcement is handled **only via HTTP request headers**.  
All data is stored **in memory** using Go maps and slices.

There is **no authentication**, **no database**, and **no external libraries**.

## Technology Stack

- Language: **Go (Golang)**
- Standard Library Only:
  - net/http
  - encoding/json
  - sync
  - context
- Storage: **In-memory (maps & slices)**
- Architecture: Clean, modular, idiomatic Go

## Role Permissions Matrix

| Action | Admin | Teacher | Student |
|------|------|--------|--------|
| Create users | Yes | No | No |
| Create courses | Yes | No | No |
| Assign teacher to course | Yes | No | No |
| Assign grades | No | Yes | No |
| View course grade summary | No | Yes | No |
| View own grades | No | No | Yes |
| View GPA | No | No | Yes |

## GPA Calculation Rules

| Score Range | Grade Point |
|------------|-------------|
| ≥ 90 | 4.0 |
| ≥ 80 | 3.0 |
| ≥ 70 | 2.0 |
| ≥ 60 | 1.0 |
| < 60 | 0.0 |

**GPA is the arithmetic mean of grade points across all courses taken by the student.**

## API Contract

### Required Headers
X-Role: Admin | Teacher | Student
X-User-ID: <integer>

Requests missing or using invalid headers are rejected.

## API Endpoints

### Health Check
GET /health

### Admin APIs

#### Create User
POST /admin/users

{
  "name": "Alice",
  "role": "Teacher"
}

#### Create User
POST /admin/courses

{
  "name": "Math"
}

#### Assign Teacher to Course
POST /admin/assign-teacher?course_id=1&teacher_id=1

### Teacher APIs

#### Assign Grade
POST /teacher/assign-grade
{
  "student_id": 2,
  "course_id": 1,
  "score": 85
}

#### View Course Grade Summary
GET /teacher/course-summary?course_id=1

### Student APIs

#### View Own Grades
GET /student/grades

#### View GPA
GET /student/gpa

{
  "student_id": 2,
  "gpa": 3.0
}

## Project Structure
```
student-grade-management-system/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── storage/
│   └── utils/
├── docs/
│   └── design.md
├── prompts.md
├── go.mod
└── README.md
```


## How to Run

##### From the project root:
go run ./cmd/server/main.go


##### Server starts at:
http://localhost:8080


## Important Notes

- All data is stored in memory and resets on server restart
- No authentication is implemented by design
- Role enforcement is explicit and transparent
- Concurrency safety ensured using mutexes
- Proper HTTP status codes used throughout

## AI Transparency

All AI prompts used during development are included **verbatim** in:
```
prompts.md
```