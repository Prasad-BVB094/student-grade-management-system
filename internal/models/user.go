package models

type Role string

const (
	RoleAdmin   Role = "Admin"
	RoleTeacher Role = "Teacher"
	RoleStudent Role = "Student"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role Role   `json:"role"`
}
