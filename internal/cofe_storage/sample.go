package cofe_storage

import "time"

// Just sample to run server

type SampleStorage struct {
}

type Config struct {
	DSN     string
	Timeout time.Duration
}

func NewCofeStorage(config Config) (storage *SampleStorage, err error) {
	return
}

func (s *SampleStorage) Close() {}

func (s *SampleStorage) GetHello() string {
	return "Hello!"
}
