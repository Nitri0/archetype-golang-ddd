package domain

type Product struct {
	ID    string
	Price int
}

func (p *Product) GetPrice() int {
	return p.Price
}
