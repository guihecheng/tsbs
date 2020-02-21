package coredata

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/timescale/tsbs/cmd/tsbs_generate_data/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_data/serialize"
)

var (
	labelEvent     = []byte("event")
	labelEventId   = []byte("eventid")
	labelPid       = []byte("pid")
	labelProtocol  = []byte("protocol")
	labelCreated   = []byte("created")
	labelName      = []byte("name")
	labelValue     = []byte("value")
	labelValuetype = []byte("valuetype")
	eventND        = common.ND(50, 10)

	eventFields = []common.LabeledDistributionMaker{
		{
			Label: labelEventId,
			DistributionMaker: func() common.Distribution {
				return common.CWD(eventND, 0, 10000000, 0)
			},
		},
		{
			Label: labelPid,
			DistributionMaker: func() common.Distribution {
				return common.CWD(eventND, 0, 100, 0)
			},
		},
		{
			Label: labelProtocol,
			DistributionMaker: func() common.Distribution {
				return common.CWD(eventND, 0, 100, 0)
			},
		},
		{
			Label: labelCreated,
			DistributionMaker: func() common.Distribution {
				return common.CWD(eventND, 0, 1000000000, 0)
			},
		},
	}
)

// EventMeasurement represents a event subset of measurements.
type EventMeasurement struct {
	*common.SubsystemMeasurement
	Name, Value, Valuetype string
}

// ToPoint serializes EventMeasurement to serialize.Point.
func (m *EventMeasurement) ToPoint(p *serialize.Point) {
	p.SetMeasurementName(labelEvent)
	copy := m.Timestamp
	p.SetTimestamp(&copy)

	p.AppendField(eventFields[0].Label, int64(m.Distributions[0].Get()))
	p.AppendField(eventFields[1].Label, int64(m.Distributions[1].Get()))
	p.AppendField(eventFields[2].Label, int64(m.Distributions[2].Get()))
	p.AppendField(eventFields[3].Label, int64(m.Distributions[3].Get()))
	p.AppendTag(labelName, m.Name)
	p.AppendTag(labelValue, m.Value)
	p.AppendTag(labelValuetype, m.Valuetype)
}

// NewEventMeasurement creates a EventMeasurement with start time.
func NewEventMeasurement(start time.Time) *EventMeasurement {
	sub := common.NewSubsystemMeasurementWithDistributionMakers(start, eventFields)
	Name := fmt.Sprintf("name_%d", rand.Intn(100000000))
	Value := fmt.Sprintf("value_%d", rand.Intn(100000000))
	Valuetype := fmt.Sprintf("valuetype_%d", rand.Intn(100))

	return &EventMeasurement{
		SubsystemMeasurement: sub,
		Name:                 Name,
		Value:                Value,
		Valuetype:            Valuetype,
	}
}
