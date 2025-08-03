package course

import (
	"Template/internal/entity"
	"Template/pkg/dbcontext"
	"context"
	"log"

	"gorm.io/gorm"
)

type repository struct {
	db     *dbcontext.DB
	logger *log.Logger
}

type Repository interface {
	Create(ctx context.Context, course entity.Course) error
	Query(ctx context.Context, page, limit int, facultyId, courseId *int, isDeleted *bool) ([]entity.Course, error)
	Update(ctx context.Context, course entity.Course) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id int) (entity.Course, error)
	GetFaculty(ctx context.Context, id string) (entity.Faculty, error)
}

func NewRepository(db *dbcontext.DB) Repository {
	return repository{db: db}
}

func (r repository) Create(ctx context.Context, course entity.Course) error {
	tx := r.db.With(ctx).Create(&course)
	return tx.Error
}

func (r repository) Query(ctx context.Context, page, limit int, facultyId, courseId *int, isDeleted *bool) ([]entity.Course, error) {
	var courses []entity.Course
	tx := r.db.With(ctx).Preload("Faculty")
	if facultyId != nil {
		tx = tx.Where("faculty_id = ?", *facultyId)
	}
	if courseId != nil {
		tx = tx.Where("id = ?", *courseId)
	}
	if isDeleted != nil && *isDeleted {
		tx = tx.Where("deleted_at IS NULL")
	}
	tx = tx.Scopes(Paginate(page, limit)).Find(&courses)
	return courses, tx.Error

}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func (r repository) Update(ctx context.Context, Course entity.Course) error {
	tx := r.db.With(ctx).Save(&Course)
	return tx.Error
}

func (r repository) Delete(ctx context.Context, id string) error {
	tx := r.db.With(ctx).Delete(&entity.Course{}, id)
	return tx.Error
}

func (r repository) Get(ctx context.Context, id int) (entity.Course, error) {
	var course entity.Course
	tx := r.db.With(ctx).Preload("Faculty").Where("id = ?", id).First(&course)
	return course, tx.Error
}

func (r repository) GetFaculty(ctx context.Context, id string) (entity.Faculty, error) {
	var falcuty entity.Faculty
	tx := r.db.With(ctx).Where("id = ?", id).First(&falcuty)
	return falcuty, tx.Error
}
