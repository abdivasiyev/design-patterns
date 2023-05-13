package main

import (
	"fmt"
	"github.com/abdivasiyev/design-patterns/behavioral/observer"
)

func main() {
	temperatureObservable := observer.NewTemperatureObservable()
	mobileObserver := observer.NewMobileObserver("mobile-observer", temperatureObservable)

	_ = temperatureObservable.Register(mobileObserver)

	temperature := mobileObserver.GetTemperature()
	fmt.Println("Initial temperature:", temperature)
	for {
		newTemperature := mobileObserver.GetTemperature()
		if temperature != newTemperature {
			fmt.Println("Temperature updated:", newTemperature)
			temperature = newTemperature
		}
	}
}
