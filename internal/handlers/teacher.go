package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"student-grade-management-system/internal/models"
	"student-grade-management-system/internal/storage"
)

type TeacherHandler struct {
	Store *storage.MemoryStore
}

/* -------- Assign Grade -------- */

func (h *TeacherHandler) AssignGrade(w http.ResponseWriter, r *http.Request) {
	var req struct {
		StudentID int `json:"student_id"`
		CourseID  int `json:"course_id"`
		Score     int `json:"score"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.StudentID <= 0 || req.CourseID <= 0 || req.Score < 0 || req.Score > 100 {
		http.Error(w, "invalid input values", http.StatusBadRequest)
		return
	}

	course, ok := h.Store.GetCourse(req.CourseID)
	if !ok {
		http.Error(w, "course not found", http.StatusBadRequest)
		return
	}

	teacherIDStr := r.Header.Get("X-User-ID")
	teacherID, _ := strconv.Atoi(teacherIDStr)

	if course.TeacherID != teacherID {
		http.Error(w, "not assigned to this course", http.StatusForbidden)
		return
	}

	student, ok := h.Store.GetUser(req.StudentID)
	if !ok || student.Role != models.RoleStudent {
		http.Error(w, "student not found", http.StatusBadRequest)
		return
	}

	h.Store.AssignGrade(req.StudentID, req.CourseID, req.Score)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("grade assigned"))
}

/* -------- Course Grade Summary -------- */

func (h *TeacherHandler) CourseSummary(w http.ResponseWriter, r *http.Request) {
	courseIDStr := r.URL.Query().Get("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil || courseID <= 0 {
		http.Error(w, "invalid course_id", http.StatusBadRequest)
		return
	}

	course, ok := h.Store.GetCourse(courseID)
	if !ok {
		http.Error(w, "course not found", http.StatusBadRequest)
		return
	}

	teacherIDStr := r.Header.Get("X-User-ID")
	teacherID, _ := strconv.Atoi(teacherIDStr)

	if course.TeacherID != teacherID {
		http.Error(w, "not assigned to this course", http.StatusForbidden)
		return
	}

	grades := h.Store.GetGradesByCourse(courseID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grades)
}
