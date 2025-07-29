package faculty

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, input CreateFacultyRequest) entity.Result
	Query(ctx context.Context) entity.Result
	Update(ctx context.Context, id string, input UpdateFacultyRequest) entity.Result
	Delete(ctx context.Context, id string) entity.Result
}

type Faculty struct {
	entity.Faculty
}

type CreateFacultyRequest struct {
	Id int
	entity.LocalizedName
}

type UpdateFacultyRequest struct {
	Id            int
	LocalizedName entity.LocalizedName
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Create(ctx context.Context, faculty CreateFacultyRequest) entity.Result {
	err := s.repo.Create(ctx, entity.Faculty{
		Name:    faculty.LocalizedName.Vi,
		EngName: faculty.LocalizedName.En,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "faculties_name_eng_key":
				return entity.Fail("DUPLICATE_FACULTY_NAME", "Tên khoa 'EN' đã tồn tại.", nil)
			case "faculties_name_key":
				return entity.Fail("DUPLICATE_FACULTY_NAME", "Tên khoa 'VI' đã tồn tại.", nil)
			}
		}
		return entity.Fail("ADD_FACULTY_FAILED", "Thêm khoa thất bại.", nil)
	}
	return entity.Ok(faculty, "")
}

func (s service) Query(ctx context.Context) entity.Result {
	faculties, err := s.repo.Query(ctx)
	if err != nil {
		return entity.Fail("GET_FACULTIES_FAILED", err.Error(), nil)
	}
	return entity.Ok(faculties, "")
}

func (s service) Update(ctx context.Context, id string, faculty UpdateFacultyRequest) entity.Result {
	if _, err := s.repo.Get(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("ADD_FACULTY_FAILED", "Khoa không tồn tại.", nil)
	}
	err := s.repo.Update(ctx, entity.Faculty{
		Name:    faculty.LocalizedName.Vi,
		EngName: faculty.LocalizedName.En,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "faculties_name_eng_key":
				return entity.Fail("DUPLICATE_FACULTY_NAME", "Tên khoa 'EN' đã tồn tại.", nil)
			case "faculties_name_key":
				return entity.Fail("DUPLICATE_FACULTY_NAME", "Tên khoa 'VI' đã tồn tại.", nil)
			}
		}
		return entity.Fail("ADD_FACULTY_FAILED", "Thêm khoa thất bại.", nil)
	}
	return entity.Ok(faculty, "")
}

func (s service) Delete(ctx context.Context, id string) entity.Result {
	err := s.repo.Delete(ctx, id)
	if err != nil {	
		return entity.Fail("DELETE_FACULTY_FAILED", err.Error(), nil)
	}
	return entity.Ok(id, "")
}
