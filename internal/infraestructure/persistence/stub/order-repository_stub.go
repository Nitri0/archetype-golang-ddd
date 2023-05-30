package stub_persistence

func NewStubOrderRepository() *StubOrderRepository {
	return &StubOrderRepository{}
}

type StubOrderRepository struct{}
