package application

import (
	"fmt"
	stub_persistence "go-archtype-ddd/internal/infraestructure/persistence/stub"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {

	t.Run("create new order", func(t *testing.T) {

		orderRepository := stub_persistence.NewStubOrderRepository()
		productRepository := stub_persistence.NewStubProductRepository()

		command := CreateOrderCommand{
			ProductIds: []string{"111", "222", "333"},
			ClientRut:  "1213114-1",
			ClientName: "Nombre Prueba",
		}

		useCase := NewCreateOrderCommandHandler(
			orderRepository,
			productRepository,
		)

		err := useCase.Exec(command)

		assert.ErrorIs(t, fmt.Errorf("product 111 not exist"), err)
	})
}
