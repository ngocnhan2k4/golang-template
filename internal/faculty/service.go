package faculty

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"context"
	"errors"
	"strconv"

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
	Id                   string `json:"id"`
	entity.LocalizedName `json:"name"`
}

type UpdateFacultyRequest struct {
	Id            string               `json:"id"`
	LocalizedName entity.LocalizedName `json:"name"`
}

type GetacultyRequest struct {
	Id            int                  `json:"id"`
	LocalizedName entity.LocalizedName `json:"name"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Create(ctx context.Context, faculty CreateFacultyRequest) entity.Result {
	id, _ := strconv.Atoi(faculty.Id)
	err := s.repo.Create(ctx, entity.Faculty{
		ID:      id,
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
	return entity.Ok(faculty, nil)
}

func (s service) Query(ctx context.Context) entity.Result {
	faculties, err := s.repo.Query(ctx)
	if err != nil {
		return entity.Fail("GET_FACULTIES_FAILED", err.Error(), nil)
	}
	res := make([]GetacultyRequest, len(faculties))
	for i := range faculties {
		res[i] = GetacultyRequest{
			Id: faculties[i].ID,
			LocalizedName: entity.LocalizedName{
				Vi: faculties[i].Name,
				En: faculties[i].EngName,
			},
		}
	}
	return entity.Ok(res, nil)
}

func (s service) Update(ctx context.Context, id string, faculty UpdateFacultyRequest) entity.Result {
	if _, err := s.repo.Get(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("ADD_FACULTY_FAILED", "Khoa không tồn tại.", nil)
	}
	intID, _ := strconv.Atoi(id)
	err := s.repo.Update(ctx, entity.Faculty{
		ID:      intID,
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
	return entity.Ok(faculty, nil)
}

func (s service) Delete(ctx context.Context, id string) entity.Result {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return entity.Fail("DELETE_FACULTY_FAILED", err.Error(), nil)
	}
	return entity.Ok(id, nil)
}
