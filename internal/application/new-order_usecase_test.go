package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {

	t.Run("create new order", func(t *testing.T) {

		orderRepository := NewStubOrderRepository()
		productRepository := NewStubProductRepository()

		command := CreateOrderCommand{}

		useCase := NewCreateOrderCommandHandler(
			orderRepository,
			productRepository,
		)

		err := useCase.Exec(command)

		assert.NoError(t, err)
	})
}
