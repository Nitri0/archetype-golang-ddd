package stub_persistence

import (
	"go-archtype-ddd/internal/domain"
)

func NewStubProductRepository() *StubProductRepository {
	return &StubProductRepository{}
}

type StubProductRepository struct{}

func (r *StubProductRepository) SearchByIds(id []string) (products []domain.Product, err error) {
	return []domain.Product{}, domain.ErrProductNotFound
}
