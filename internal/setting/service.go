package setting

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"context"
)

type Service interface {
	Get(ctx context.Context) GetSettingRequest
	Update(ctx context.Context, input string) bool
}

type GetSettingRequest struct {
	Domain string `json:"domain"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Update(ctx context.Context, domain string) bool {
	if len(domain) == 0 {
		return false
	}
	err := s.repo.Update(ctx, entity.Setting{
		EmailDomain: domain,
	})
	return err == nil
}

func (s service) Get(ctx context.Context) GetSettingRequest {
	setting, _ := s.repo.Get(ctx)

	res := GetSettingRequest{
		Domain: setting.EmailDomain,
	}
	return res
}
