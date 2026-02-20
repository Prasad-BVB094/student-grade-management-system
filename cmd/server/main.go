package main

import (
	"log"
	"net/http"

	"student-grade-management-system/internal/handlers"
	"student-grade-management-system/internal/middleware"
	"student-grade-management-system/internal/models"
	"student-grade-management-system/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()

	adminHandler := &handlers.AdminHandler{Store: store}
	teacherHandler := &handlers.TeacherHandler{Store: store}
	studentHandler := &handlers.StudentHandler{Store: store}

	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Admin
	mux.Handle("/admin/users",
		middleware.RoleMiddleware(http.HandlerFunc(adminHandler.CreateUser), models.RoleAdmin),
	)
	mux.Handle("/admin/courses",
		middleware.RoleMiddleware(http.HandlerFunc(adminHandler.CreateCourse), models.RoleAdmin),
	)
	mux.Handle("/admin/assign-teacher",
		middleware.RoleMiddleware(http.HandlerFunc(adminHandler.AssignTeacher), models.RoleAdmin),
	)

	// Teacher
	mux.Handle("/teacher/assign-grade",
		middleware.RoleMiddleware(http.HandlerFunc(teacherHandler.AssignGrade), models.RoleTeacher),
	)
	mux.Handle("/teacher/course-summary",
		middleware.RoleMiddleware(http.HandlerFunc(teacherHandler.CourseSummary), models.RoleTeacher),
	)

	// Student
	mux.Handle("/student/grades",
		middleware.RoleMiddleware(http.HandlerFunc(studentHandler.ViewGrades), models.RoleStudent),
	)
	mux.Handle("/student/gpa",
		middleware.RoleMiddleware(http.HandlerFunc(studentHandler.ViewGPA), models.RoleStudent),
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server started on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
