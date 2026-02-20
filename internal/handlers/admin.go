package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"student-grade-management-system/internal/models"
	"student-grade-management-system/internal/storage"
)

type AdminHandler struct {
	Store *storage.MemoryStore
}

/* -------- Create User -------- */

func (h *AdminHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string      `json:"name"`
		Role models.Role `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || (req.Role != models.RoleStudent && req.Role != models.RoleTeacher) {
		http.Error(w, "invalid name or role", http.StatusBadRequest)
		return
	}

	user := h.Store.CreateUser(req.Name, req.Role)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

/* -------- Create Course -------- */

func (h *AdminHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	course := h.Store.CreateCourse(req.Name, 0)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

/* -------- Assign Teacher to Course -------- */

func (h *AdminHandler) AssignTeacher(w http.ResponseWriter, r *http.Request) {
	courseIDStr := r.URL.Query().Get("course_id")
	teacherIDStr := r.URL.Query().Get("teacher_id")

	courseID, err1 := strconv.Atoi(courseIDStr)
	teacherID, err2 := strconv.Atoi(teacherIDStr)

	if err1 != nil || err2 != nil {
		http.Error(w, "invalid course_id or teacher_id", http.StatusBadRequest)
		return
	}

	teacher, ok := h.Store.GetUser(teacherID)
	if !ok || teacher.Role != models.RoleTeacher {
		http.Error(w, "teacher not found", http.StatusBadRequest)
		return
	}

	course, ok := h.Store.GetCourse(courseID)
	if !ok {
		http.Error(w, "course not found", http.StatusBadRequest)
		return
	}

	course.TeacherID = teacherID
	h.Store.Courses[course.ID] = course

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(course)
}
