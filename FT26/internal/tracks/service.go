package tracks

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]Circuit, error)
	GetByID(ctx context.Context, id string) (*Circuit, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(ctx context.Context) ([]Circuit, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id string) (*Circuit, error) {
	return s.repo.GetByID(ctx, id)
}
