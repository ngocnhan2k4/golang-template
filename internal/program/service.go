package program

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
	Create(ctx context.Context, input CreateProgramRequest) entity.Result
	Query(ctx context.Context) entity.Result
	Update(ctx context.Context, id string, input UpdateProgramRequest) entity.Result
	Delete(ctx context.Context, id string) entity.Result
}

type Program struct {
	entity.Program
}

type CreateProgramRequest struct {
	Id                   string `json:"id"`
	entity.LocalizedName `json:"name"`
}

type UpdateProgramRequest struct {
	Id                   string `json:"id"`
	entity.LocalizedName `json:"name"`
}

type GetProgramRequest struct {
	Id            string               `json:"id"`
	LocalizedName entity.LocalizedName `json:"name"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Create(ctx context.Context, program CreateProgramRequest) entity.Result {
	err := s.repo.Create(ctx, entity.Program{
		Name:    program.LocalizedName.Vi,
		EngName: program.LocalizedName.En,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "programs_name_eng_key":
				return entity.Fail("DUPLICATE_PROGRAM_NAME", "Tên chương trình 'EN' đã tồn tại.", nil)
			case "programs_name_key":
				return entity.Fail("DUPLICATE_PROGRAM_NAME", "Tên chương trình 'VI' đã tồn tại.", nil)
			}
		}
		return entity.Fail("ADD_PROGRAM_FAILED", "Thêm chương trình thất bại.", nil)
	}
	return entity.Ok(program, nil)
}

func (s service) Query(ctx context.Context) entity.Result {
	programs, err := s.repo.Query(ctx)
	if err != nil {
		return entity.Fail("GET_PROGRAMS_FAILED", err.Error(), nil)
	}
	res := make([]GetProgramRequest, len(programs))
	for i := range programs {
		id := strconv.Itoa(programs[i].ID)
		res[i] = GetProgramRequest{
			Id: id,
			LocalizedName: entity.LocalizedName{
				Vi: programs[i].Name,
				En: programs[i].EngName,
			},
		}
	}
	return entity.Ok(res, nil)
}

func (s service) Update(ctx context.Context, id string, Program UpdateProgramRequest) entity.Result {
	if _, err := s.repo.Get(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("PROGRAM_NOT_FOUND", "Chương trình không tồn tại.", nil)
	}
	intID, _ := strconv.Atoi(id)
	err := s.repo.Update(ctx, entity.Program{
		ID:      intID,
		Name:    Program.LocalizedName.Vi,
		EngName: Program.LocalizedName.En,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "programs_name_eng_key":
				return entity.Fail("DUPLICATE_PROGRAM_NAME", "Tên chương trình 'EN' đã tồn tại.", nil)
			case "programs_name_key":
				return entity.Fail("DUPLICATE_PROGRAM_NAME", "Tên chương trình 'VI' đã tồn tại.", nil)
			}
		}
		return entity.Fail("UPDATE_PROGRAM_FAILED", "Cập nhật chương trình thất bại.", nil)
	}
	return entity.Ok(Program, nil)
}

func (s service) Delete(ctx context.Context, id string) entity.Result {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return entity.Fail("DELETE_PROGRAM_FAILED", err.Error(), nil)
	}
	return entity.Ok(id, nil)
}
