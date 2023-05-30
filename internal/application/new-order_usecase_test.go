package application

import (
	stub_persistence "go-archtype-ddd/internal/infraestructure/persistence/stub"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {

	t.Run("create new order", func(t *testing.T) {

		orderRepository := stub_persistence.NewStubOrderRepository()
		productRepository := stub_persistence.NewStubProductRepository()

		command := CreateOrderCommand{}

		useCase := NewCreateOrderCommandHandler(
			orderRepository,
			productRepository,
		)

		err := useCase.Exec(command)

		assert.NoError(t, err)
	})
}
