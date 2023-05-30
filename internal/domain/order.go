package domain

type Order struct {
	ID         string
	Products   []Product
	ClientRut  string
	ClientName string
}

func (o Order) GetTotalPrice() (total int) {
	for _, p := range o.Products {
		total += p.GetPrice()
	}
	return total
}
