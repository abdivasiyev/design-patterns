package observer

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

type Observable interface {
	Register(observer Observer) error
	UnRegister(observer Observer) error
	Notify(ctx context.Context) error
}

type TemperatureObservable struct {
	observers   map[string]Observer
	temperature int
}

func (t *TemperatureObservable) GetTemperature() int {
	return t.temperature
}

var (
	ErrObserverAlreadyRegistered = errors.New("observer already registered")
	ErrObserverIsNotExists       = errors.New("observer is not exists")
)

func (t *TemperatureObservable) Register(observer Observer) error {
	id := observer.ID()
	if _, ok := t.observers[id]; ok {
		return ErrObserverAlreadyRegistered
	}

	t.observers[id] = observer
	return nil
}

func (t *TemperatureObservable) UnRegister(observer Observer) error {
	id := observer.ID()
	if _, ok := t.observers[id]; !ok {
		return ErrObserverIsNotExists
	}
	delete(t.observers, id)
	return nil
}

func (t *TemperatureObservable) Notify(ctx context.Context) error {
	for _, o := range t.observers {
		if err := o.Update(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (t *TemperatureObservable) generateTemperature() {
	for {
		select {
		case <-time.After(3 * time.Second):
			t.temperature = rand.Intn(100)
			_ = t.Notify(context.Background())
		}
	}
}

func NewTemperatureObservable() *TemperatureObservable {
	t := &TemperatureObservable{
		observers: make(map[string]Observer),
	}

	go t.generateTemperature()

	return t
}
