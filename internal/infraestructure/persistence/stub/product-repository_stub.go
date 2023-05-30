package stub_persistence

func NewStubProductRepository() *StubProductRepository {
	return &StubProductRepository{}
}

type StubProductRepository struct{}
