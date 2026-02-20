package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"student-grade-management-system/internal/models"
	"student-grade-management-system/internal/storage"
	"student-grade-management-system/internal/utils"
)

type StudentHandler struct {
	Store *storage.MemoryStore
}

/* -------- View Own Grades -------- */

func (h *StudentHandler) ViewGrades(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, ok := h.Store.GetUser(userID)
	if !ok || user.Role != models.RoleStudent {
		http.Error(w, "student not found", http.StatusForbidden)
		return
	}

	grades := h.Store.GetGradesByStudent(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grades)
}

/* -------- View GPA -------- */

func (h *StudentHandler) ViewGPA(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, ok := h.Store.GetUser(userID)
	if !ok || user.Role != models.RoleStudent {
		http.Error(w, "student not found", http.StatusForbidden)
		return
	}

	grades := h.Store.GetGradesByStudent(userID)

	var scores []int
	for _, g := range grades {
		scores = append(scores, g.Score)
	}

	gpa := utils.CalculateGPA(scores)

	resp := struct {
		StudentID int     `json:"student_id"`
		GPA       float64 `json:"gpa"`
	}{
		StudentID: userID,
		GPA:       gpa,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
