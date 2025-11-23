package products

import (
	"context"

	repo "github.com/FelipePn10/ecomms/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	repository repo.Querier
}

func NewService(repository repo.Querier) Service {
	return &svc{repository: repository}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repository.ListProducts(ctx)

}
