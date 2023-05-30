package application

import (
	"go-archtype-ddd/internal/domain"
)

func NewCreateOrderCommandHandler(
	OrderRepository domain.OrderRepository,
	ProductRepository domain.ProductRepository) *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{
		OrderRepository:   OrderRepository,
		ProductRepository: ProductRepository,
	}
}

type ICreateOrderCommandHandler interface {
	Exec(request CreateOrderCommand) (err error)
}

type CreateOrderCommand struct {
	ProductIds []string
	ClientName string
	ClientRut  string
}

type CreateOrderCommandHandler struct {
	OrderRepository   domain.OrderRepository
	ProductRepository domain.ProductRepository
}

func (a *CreateOrderCommandHandler) Exec(command CreateOrderCommand) (order domain.Order, err error) {
	products, err := a.ProductRepository.SearchByIds(command.ProductIds)

	if err != nil {
		return
	}

	order = domain.Order{
		ID:         "",
		Products:   products,
		ClientRut:  command.ClientRut,
		ClientName: command.ClientName,
	}

	return
}
