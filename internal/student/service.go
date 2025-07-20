package student

import(
	"Template/internal/entity"
	"context"
)

// Service encapsulates usecase logic for students.
type Service interface {
	Query(ctx context.Context) ([]entity.Student, error)
}

// Student represents the data about an student.
type Student struct{
	entity.Student
}

type service struct {
	repo Repository
}

func (s service) Query(ctx context.Context) ([]entity.Student, error) {
	return s.repo.Query(ctx)	
}

// NewService creates a new student service.
func NewService (repo Repository) Service {
	return service{repo : repo}
}