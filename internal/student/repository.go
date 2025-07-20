package student

import (
	"Template/internal/entity"
	"Template/pkg/dbcontext"
	"context"
)

// Repository encapsulates the logic to access students from the data source.
type Repository interface {
	Query(ctx context.Context) ([]entity.Student, error)
}

// repository persists students in database
type repository struct{
	db *dbcontext.DB
}

// NewRepository creates a new student repository.
func NewRepository(db *dbcontext.DB) Repository{
	return repository{db: db}
}


func (r repository) Query(ctx context.Context) ([]entity.Student, error){
	var students []entity.Student
	result := r.db.With(ctx).Find(&students)

	return students, result.Error
}