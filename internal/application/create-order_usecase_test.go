package application

import (
	"archetype-golang-ddd/internal/domain"
	stub_persistence "archetype-golang-ddd/internal/infraestructure/persistence/stub"
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

		expectedProducts := []domain.Product{
			{ID: "111", Price: 100},
			{ID: "222", Price: 50},
			{ID: "333", Price: 50},
		}

		expectedOrder := domain.Order{
			ID:         "",
			Products:   expectedProducts,
			ClientRut:  "1213114-1",
			ClientName: "Nombre Prueba",
		}

		productRepository.On("SearchByIds", []string{"111", "222", "333"}).Return(expectedProducts, nil)

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
		assert.Equal(t, 200, order.GetTotalPrice())
		assert.Equal(t, expectedOrder, order)
	})
}
