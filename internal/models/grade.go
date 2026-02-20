package models

type Grade struct {
	StudentID int `json:"student_id"`
	CourseID  int `json:"course_id"`
	Score     int `json:"score"`
}
