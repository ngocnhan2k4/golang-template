package program

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
	Create(ctx context.Context, program entity.Program) error
	Query(ctx context.Context) ([]entity.Program, error)
	Update(ctx context.Context, program entity.Program) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (entity.Program, error)
}

func NewRepository(db *dbcontext.DB) Repository {
	return repository{db: db}
}

func (r repository) Create(ctx context.Context, program entity.Program) error {
	tx := r.db.With(ctx).Create(&program)
	return tx.Error
}

func (r repository) Query(ctx context.Context) ([]entity.Program, error) {
	var programs []entity.Program
	tx := r.db.With(ctx).Find(&programs)
	return programs, tx.Error

}

func (r repository) Update(ctx context.Context, program entity.Program) error {
	tx := r.db.With(ctx).Save(&program)
	return tx.Error
}

func (r repository) Delete(ctx context.Context, id string) error {
	tx := r.db.With(ctx).Delete(&entity.Program{}, id)
	return tx.Error
}

func (r repository) Get(ctx context.Context, id string) (entity.Program, error) {
	var program entity.Program
	tx := r.db.With(ctx).Where("id = ?", id).First(&program)
	return program, tx.Error
}
