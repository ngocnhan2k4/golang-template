package status

import (
	"Template/internal/entity"
	"Template/pkg/dbcontext"
	"context"
	"log"
)


type repository struct{
	db *dbcontext.DB
	logger *log.Logger
}

type Repository interface{
	Create(ctx context.Context, status entity.Status) error
	Query(ctx context.Context) ([]entity.Status, error)
	Update(ctx context.Context, status entity.Status) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (entity.Status, error)
}

func NewRepository(db * dbcontext.DB) Repository{
	return repository{db : db}
}

func (r repository) Create(ctx context.Context, status entity.Status) error{
	tx := r.db.With(ctx).Create(&status)
	return tx.Error
}

func (r repository) Query(ctx context.Context) ([]entity.Status, error){
	var statuses []entity.Status
	tx := r.db.With(ctx).Find(&statuses)
	return statuses, tx.Error

}

func (r repository) Update(ctx context.Context, status entity.Status) error{
	tx := r.db.With(ctx).Save(&status)
	return tx.Error
}


func (r repository) Delete(ctx context.Context, id string) error{
	tx := r.db.With(ctx).Delete(&entity.Status{}, id)
	return tx.Error
}

func (r repository) Get(ctx context.Context, id string) (entity.Status,error){
	var status entity.Status
	tx := r.db.With(ctx).Where("id = ?", id).First(&status)
	return status, tx.Error
}