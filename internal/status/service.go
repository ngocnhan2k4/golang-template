package status

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
	Create(ctx context.Context, input CreateStatusRequest) entity.Result
	Query(ctx context.Context) entity.Result
	Update(ctx context.Context, id string, input UpdateStatusRequest) entity.Result
	Delete(ctx context.Context, id string) entity.Result
}

type Status struct {
	entity.Status
}

type CreateStatusRequest struct {
	Id                   string `json:"id"`
	entity.LocalizedName `json:"name"`
	Order                int `json:"order"`
}

type UpdateStatusRequest struct {
	Id                   string `json:"id"`
	entity.LocalizedName `json:"name"`
	Order                int `json:"order"`
}

type GetStatusRequest struct {
	Id            string                  `json:"id"`
	LocalizedName entity.LocalizedName `json:"name"`
	Order         int                  `json:"order"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Create(ctx context.Context, status CreateStatusRequest) entity.Result {
	err := s.repo.Create(ctx, entity.Status{
		Name:    status.LocalizedName.Vi,
		EngName: status.LocalizedName.En,
		Order:   status.Order,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "student_statuses_name_eng_key":
				return entity.Fail("DUPLICATE_STUDENT_STATUS_NAME", "Tên trạng thái 'EN' đã tồn tại.", nil)
			case "student_statuses_name_key":
				return entity.Fail("DUPLICATE_STUDENT_STATUS_NAME", "Tên trạng thái 'VI' đã tồn tại.", nil)
			}
		}
		return entity.Fail("ADD_STUDENT_STATUS_FAILED", "Thêm trạng thái thất bại.", nil)
	}
	return entity.Ok(status, nil)
}

func (s service) Query(ctx context.Context) entity.Result {
	statuses, err := s.repo.Query(ctx)
	if err != nil {
		return entity.Fail("GET_STUDENT_STATUSES_FAILED", err.Error(), nil)
	}
	res := make([]GetStatusRequest, len(statuses))
	for i := range statuses {
		res[i] = GetStatusRequest{
			Id: strconv.Itoa(statuses[i].ID),
			LocalizedName: entity.LocalizedName{
				Vi: statuses[i].Name,
				En: statuses[i].EngName,
			},
			Order: statuses[i].Order,
		}
	}
	return entity.Ok(res, nil)
}

func (s service) Update(ctx context.Context, id string, Status UpdateStatusRequest) entity.Result {
	if _, err := s.repo.Get(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Fail("STUDENT_STATUS_NOT_FOUND", "Trạng thái sinh viên không tồn tại.", nil)
	}
	intID, _ := strconv.Atoi(id)
	err := s.repo.Update(ctx, entity.Status{
		ID:      intID,
		Name:    Status.LocalizedName.Vi,
		EngName: Status.LocalizedName.En,
		Order:   Status.Order,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "student_statuses_name_eng_key":
				return entity.Fail("DUPLICATE_STUDENT_STATUS_NAME", "Tên trạng thái 'EN' đã tồn tại.", nil)
			case "student_statuses_name_key":
				return entity.Fail("DUPLICATE_STUDENT_STATUS_NAME", "Tên trạng thái 'VI' đã tồn tại.", nil)
			}
		}
		return entity.Fail("UPDATE_STUDENT_STATUS_FAILED", "Cập nhật trạng thái thất bại.", nil)
	}
	return entity.Ok(Status, nil)
}

func (s service) Delete(ctx context.Context, id string) entity.Result {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return entity.Fail("DELETE_STUDENT_STATUS_FAILED", err.Error(), nil)
	}
	return entity.Ok(id, nil)
}
