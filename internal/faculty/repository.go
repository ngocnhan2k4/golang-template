package faculty

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
	Create(ctx context.Context, falcuty entity.Faculty) error
	Query(ctx context.Context) ([]entity.Faculty, error)
	Update(ctx context.Context, faculty entity.Faculty) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (entity.Faculty, error)
}

func NewRepository(db * dbcontext.DB) Repository{
	return repository{db : db}
}

func (r repository) Create(ctx context.Context, falcuty entity.Faculty) error{
	tx := r.db.With(ctx).Create(&falcuty)
	return tx.Error
}

func (r repository) Query(ctx context.Context) ([]entity.Faculty, error){
	var faculties []entity.Faculty
	tx := r.db.With(ctx).Find(&faculties)
	return faculties, tx.Error

}

func (r repository) Update(ctx context.Context, faculty entity.Faculty) error{
	tx := r.db.With(ctx).Save(&faculty)
	return tx.Error
}


func (r repository) Delete(ctx context.Context, id string) error{
	tx := r.db.With(ctx).Delete(&entity.Faculty{}, id)
	return tx.Error
}

func (r repository) Get(ctx context.Context, id string) (entity.Faculty,error){
	var falcuty entity.Faculty
	tx := r.db.With(ctx).Where("id = ?", id).First(&falcuty)
	return falcuty, tx.Error
}