package cofe_storage

// Just sample to run server

type SampleStorage struct {
}

func NewCofeStorage(filename string) (storage *SampleStorage, err error) {
	return
}

func (s *SampleStorage) Close() {}

func (s *SampleStorage) GetHello() string {
	return "Hello!"
}
