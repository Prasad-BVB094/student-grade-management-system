package storage

import (
	"sync"

	"student-grade-management-system/internal/models"
)

type MemoryStore struct {
	mu sync.RWMutex

	Users   map[int]models.User
	Courses map[int]models.Course
	Grades  []models.Grade

	nextUserID   int
	nextCourseID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Users:        make(map[int]models.User),
		Courses:      make(map[int]models.Course),
		Grades:       []models.Grade{},
		nextUserID:   1,
		nextCourseID: 1,
	}
}

/* ---------- User Operations ---------- */

func (s *MemoryStore) CreateUser(name string, role models.Role) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := models.User{
		ID:   s.nextUserID,
		Name: name,
		Role: role,
	}
	s.Users[user.ID] = user
	s.nextUserID++
	return user
}

func (s *MemoryStore) GetUser(id int) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.Users[id]
	return user, ok
}

/* ---------- Course Operations ---------- */

func (s *MemoryStore) CreateCourse(name string, teacherID int) models.Course {
	s.mu.Lock()
	defer s.mu.Unlock()

	course := models.Course{
		ID:        s.nextCourseID,
		Name:      name,
		TeacherID: teacherID,
	}
	s.Courses[course.ID] = course
	s.nextCourseID++
	return course
}

func (s *MemoryStore) GetCourse(id int) (models.Course, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	course, ok := s.Courses[id]
	return course, ok
}

/* ---------- Grade Operations ---------- */

func (s *MemoryStore) AssignGrade(studentID, courseID, score int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Grades = append(s.Grades, models.Grade{
		StudentID: studentID,
		CourseID:  courseID,
		Score:     score,
	})
}

func (s *MemoryStore) GetGradesByStudent(studentID int) []models.Grade {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.Grade
	for _, g := range s.Grades {
		if g.StudentID == studentID {
			result = append(result, g)
		}
	}
	return result
}

func (s *MemoryStore) GetGradesByCourse(courseID int) []models.Grade {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.Grade
	for _, g := range s.Grades {
		if g.CourseID == courseID {
			result = append(result, g)
		}
	}
	return result
}
