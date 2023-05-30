package domain

type OrderRepository interface {
}

type ProductRepository interface {
	SearchByIds([]string) ([]Product, error)
}
