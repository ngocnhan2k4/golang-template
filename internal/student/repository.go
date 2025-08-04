package student

import (
	"Template/internal/entity"
	"Template/pkg/dbcontext"
	"context"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type repository struct {
	db     *dbcontext.DB
	logger *log.Logger
}

type Repository interface {
	Create(ctx context.Context, student entity.Student) (int, error)
	Query(ctx context.Context, page, limit int, classId, semester, year *int) ([]entity.Class, error)
	Update(ctx context.Context, class entity.Class) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id int) (entity.Class, error)
	GetCourse(ctx context.Context, id string) (entity.Course, error)
	GetEmailSetting(ctx context.Context) (entity.Setting, error)
	CheckNoExistingEmail(ctx context.Context, email string) bool
	CheckNoExistingPhone(ctx context.Context, phone string) bool
	CheckValidFaculty(ctx context.Context, facultyId string) bool
	CheckValidProgram(ctx context.Context, programId string) bool
	CheckValidStatus(ctx context.Context, statusId string) bool
	CreateIdentityDocument(ctx context.Context, identity entity.IdentityDocument) error
}

func NewRepository(db *dbcontext.DB) Repository {
	return repository{db: db}
}

func (r repository) Create(ctx context.Context, student entity.Student) (int, error) {
	tx := r.db.With(ctx).Create(&student)
	return student.Id, tx.Error
}

func (r repository) Query(ctx context.Context, page, limit int, classId, semester, year *int) ([]entity.Class, error) {
	var classes []entity.Class
	tx := r.db.With(ctx).Preload("Course")
	if classId != nil {
		tx = tx.Where("id = ?", *classId)
	}
	if semester != nil {
		tx = tx.Where("semester = ?", *semester)
	}
	if year != nil {
		tx = tx.Where("academic_year = ?", *year)
	}
	tx = tx.Scopes(Paginate(page, limit)).Find(&classes)
	return classes, tx.Error

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

func (r repository) Update(ctx context.Context, class entity.Class) error {
	tx := r.db.With(ctx).Save(&class)
	return tx.Error
}

func (r repository) Delete(ctx context.Context, id string) error {
	tx := r.db.With(ctx).Delete(&entity.Class{}, id)
	return tx.Error
}

func (r repository) Get(ctx context.Context, id int) (entity.Class, error) {
	var class entity.Class
	tx := r.db.With(ctx).Where("id = ?", id).First(&class)
	return class, tx.Error
}

func (r repository) GetCourse(ctx context.Context, id string) (entity.Course, error) {
	var course entity.Course
	intId, _ := strconv.Atoi(id)
	tx := r.db.With(ctx).Where("id = ?", intId).First(&course)
	return course, tx.Error
}

func (r repository) GetEmailSetting(ctx context.Context) (entity.Setting, error) {
	var setting entity.Setting
	tx := r.db.With(ctx).Last(&setting)
	return setting, tx.Error
}

func (r repository) CheckNoExistingEmail(ctx context.Context, email string) bool {
	var count int64
	tx := r.db.With(ctx).Model(&entity.Student{}).Where("email = ?", email).Count(&count)
	if tx.Error != nil {
		return false
	}
	return count == 0
}

func (r repository) CheckNoExistingPhone(ctx context.Context, phone string) bool {
	var count int64
	tx := r.db.With(ctx).Model(&entity.Student{}).Where("phone = ?", phone).Count(&count)
	if tx.Error != nil {
		return false
	}
	return count == 0
}

func (r repository) CheckValidFaculty(ctx context.Context, facultyId string) bool {
	var count int64
	tx := r.db.With(ctx).Model(&entity.Faculty{}).Where("id = ?", facultyId).Count(&count)
	if tx.Error != nil {
		return false
	}
	return count > 0
}

func (r repository) CheckValidProgram(ctx context.Context, programId string) bool {
	var count int64
	tx := r.db.With(ctx).Model(&entity.Program{}).Where("id = ?", programId).Count(&count)
	if tx.Error != nil {
		return false
	}
	return count > 0
}

func (r repository) CheckValidStatus(ctx context.Context, statusId string) bool {
	var count int64
	tx := r.db.With(ctx).Model(&entity.Status{}).Where("id = ?", statusId).Count(&count)
	if tx.Error != nil {
		return false
	}
	return count > 0
}

func (r repository) CreateIdentityDocument(ctx context.Context, identity entity.IdentityDocument) error {
	tx := r.db.With(ctx).Create(&identity)
	return tx.Error
}
