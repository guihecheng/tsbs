package coredata

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/timescale/tsbs/cmd/tsbs_generate_data/common"
)

const (
	deviceNameFmt = "device_%d"
)

// Device models a device outfitted with an IoT device which sends back measurements.
type Device struct {
	simulatedMeasurements []common.SimulatedMeasurement
	tags                  []common.Tag
}

// TickAll advances all Distributions of a Device.
func (t *Device) TickAll(d time.Duration) {
	for i := range t.simulatedMeasurements {
		t.simulatedMeasurements[i].Tick(d)
	}
}

// Measurements returns the devices measurements.
func (t Device) Measurements() []common.SimulatedMeasurement {
	return t.simulatedMeasurements
}

// Tags returns the device tags.
func (t Device) Tags() []common.Tag {
	return t.tags
}

func newDeviceMeasurements(start time.Time) []common.SimulatedMeasurement {
	return []common.SimulatedMeasurement{
		NewEventMeasurement(start),
	}
}

// NewDevice creates a new device in a simulated iot use case
func NewDevice(i int, start time.Time) common.Generator {
	device := newDeviceWithMeasurementGenerator(i, start, newDeviceMeasurements)
	return &device
}

func newDeviceWithMeasurementGenerator(i int, start time.Time, generator func(time.Time) []common.SimulatedMeasurement) Device {
	sm := generator(start)

	h := Device{
		tags: []common.Tag{
			{Key: []byte("name"), Value: fmt.Sprintf(deviceNameFmt, i)},
			{Key: []byte("did"), Value: rand.Intn(100000000)},
		},
		simulatedMeasurements: sm,
	}

	return h
}
