package observer

import "context"

type Observer interface {
	ID() string
	Update(ctx context.Context) error
}

type MobileObserver struct {
	id                  string
	temperatureStrategy TemperatureStrategy
	temperature         int
}

func NewMobileObserver(id string, temperatureStrategy TemperatureStrategy) *MobileObserver {
	return &MobileObserver{
		id:                  id,
		temperatureStrategy: temperatureStrategy,
	}
}

func (o *MobileObserver) ID() string {
	return o.id
}

func (o *MobileObserver) Update(ctx context.Context) error {
	o.temperature = o.temperatureStrategy.GetTemperature()
	return nil
}

func (o *MobileObserver) GetTemperature() int {
	return o.temperature
}
