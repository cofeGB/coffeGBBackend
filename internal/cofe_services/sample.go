package cofe_services

import (
	"context"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/nawmenu"
)

// Sample storage interface
// package cofe_storage should contains an implementation

type CofeStorage interface {
	GetHello() string
	ReadNawMenu(ctx context.Context) ([]nawmenu.NawMenu, error)
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
