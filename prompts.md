Implement Capstone Project 3: Student Grade Management System API.

================================================
AUTHORITATIVE PROJECT SPECIFICATION
================================================

PROJECT DESCRIPTION:
Build a REST API for managing student grades similar to university portals (e.g., Canvas, Blackboard). The system must support role-based operations, course and grade management, and GPA calculation.

This is an INDIVIDUAL project.

ROLES (EXACTLY THREE):
Admin
Teacher
Student

ADMIN RESPONSIBILITIES:
- Create users (students and teachers)
- Create courses
- Assign a teacher to a course

TEACHER RESPONSIBILITIES:
- Assign grades to students for a course
- View grade summary for a course

STUDENT RESPONSIBILITIES:
- View their own grades
- View their GPA

================================================
GPA CALCULATION RULES (MANDATORY)
================================================
- Score >= 90  → 4.0
- Score >= 80  → 3.0
- Score >= 70  → 2.0
- Score >= 60  → 1.0
- Else         → 0.0

GPA is the arithmetic mean of grade points across all courses taken by the student.

================================================
TECHNICAL CONSTRAINTS
================================================
- Language: Go (Golang)
- Backend-only REST API
- Use only Go standard library (net/http, encoding/json, etc.)
- No external frameworks or libraries
- No authentication mechanisms (JWT, OAuth, sessions, etc.)
- Role enforcement via HTTP request headers only
- In-memory data storage only (maps/slices)
- Clean, modular, idiomatic Go project structure
- Proper HTTP status codes
- Proper error handling (no panics)

================================================
EXECUTION REQUIREMENTS (CRITICAL)
================================================
- Code must be 100% correct
- Code must compile without errors
- Code must run without runtime failures
- No missing imports, undefined identifiers, or package mismatches
- All files must be internally consistent
- Project must run using:
  go run ./cmd/server/main.go

Assume zero human code review or correction.

================================================
MANDATORY REPOSITORY REQUIREMENTS
================================================

- Complete source code
- README.md explaining design and implementation
- A markdown design document (combined or separate)
- All AI prompts used, included verbatim

================================================
README GENERATION RULE
================================================

- README.md must be generated LAST
- README.md must be generated ONLY when explicitly instructed with:
  "generate README now"

================================================
AI TRANSPARENCY RULE
================================================

- All prompts used must be included verbatim
- Prompts must not be summarized or rewritten