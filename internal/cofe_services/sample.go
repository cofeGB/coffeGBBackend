package cofe_services

// Sample storage interface
// package cofe_storage should contains an implementation

type CofeStorage interface {
	GetHello() string
}

// Just sample service to start app

type CofeService struct {
	storage CofeStorage
}

func NewCofeService(storage CofeStorage) (svc *CofeService) {
	return &CofeService{storage: storage}
}

func (s *CofeService) Hello() string {
	return s.storage.GetHello() // service uses storage methods
}
