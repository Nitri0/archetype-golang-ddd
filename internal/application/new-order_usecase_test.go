package application

import (
	"go-archtype-ddd/internal/domain"
	stub_persistence "go-archtype-ddd/internal/infraestructure/persistence/stub"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) SearchByIds(id []string) (products []domain.Product, err error) {
	args := m.Called(id)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func TestNewOrder(t *testing.T) {

	t.Run("create new order", func(t *testing.T) {

		orderRepository := stub_persistence.NewStubOrderRepository()
		productRepository := new(ProductRepositoryMock)

		expectedOrder := &Order{
			ID:         "",
			ProductIds: []string{"111", "222", "333"},
			ClientRut:  "1213114-1",
			ClientName: "Nombre Prueba",
		}
		productRepository.On("SearchByIds", 1).Return(expectedOrder, nil)

		command := CreateOrderCommand{
			ProductIds: []string{"111", "222", "333"},
			ClientRut:  "1213114-1",
			ClientName: "Nombre Prueba",
		}

		useCase := NewCreateOrderCommandHandler(
			orderRepository,
			productRepository,
		)

		order, err := useCase.Exec(command)

		assert.NoError(t, err)
		productRepository.AssertExpectations(t)
		assert.Equal(t, expectedOrder, order)
	})
}
