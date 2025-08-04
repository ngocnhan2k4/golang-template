package setting

import (
	"Template/internal/entity"
	"Template/pkg/dbcontext"
	"context"
	"log"
)

type repository struct {
	db     *dbcontext.DB
	logger *log.Logger
}

type Repository interface {
	Update(ctx context.Context, setting entity.Setting) error
	Get(ctx context.Context) (entity.Setting, error)
}

func NewRepository(db *dbcontext.DB) Repository {
	return repository{db: db}
}

func (r repository) Update(ctx context.Context, setting entity.Setting) error {
	setting.EmailPattern = "^([\\w\\.\\-]+)@([\\w\\-]+)((\\.[\\w]{2,3})+)$"

	tx := r.db.With(ctx).Save(&setting)
	return tx.Error
}

func (r repository) Get(ctx context.Context) (entity.Setting, error) {
	var setting entity.Setting
	tx := r.db.With(ctx).Last(&setting)
	return setting, tx.Error
}
