package singleton

// Singleton interface
type Singleton interface {
	AddOne() int
}

type singleton struct {
	counter int
}

func (s *singleton) AddOne() int {
	s.counter++
	return s.counter
}

var instance *singleton

// GetInstance ...
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}
