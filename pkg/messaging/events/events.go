package events

import (
	"time"

	"github.com/diwise/iot-core/pkg/messaging/topics"
	"github.com/farshidtz/senml/v2"	
)

type MessageAccepted struct {
	Sensor    string     `json:"sensorID"`
	Pack      senml.Pack `json:"pack"`
	Timestamp string     `json:"timestamp"`
}

func NewMessageAccepted(sensor string, pack senml.Pack) *MessageAccepted {
	
	msg := &MessageAccepted{
		Sensor:    sensor,
		Pack:      pack,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	return msg
}

func (m *MessageAccepted) ContentType() string {
	return "application/json"
}

func (m *MessageAccepted) TopicName() string {
	return topics.MessageAccepted
}

func (m MessageAccepted) AtLocation(latitude, longitude float64) MessageAccepted {
	if m.IsLocated() {
		return m
	}

	if m.Latitude() == 0 {
		lat := &senml.Record{
			Unit:  senml.UnitLat,
			Value: &latitude,
		}
		m.Pack = append(m.Pack, *lat)
	}

	if m.Longitude() == 0 {
		lon := &senml.Record{
			Unit:  senml.UnitLon,
			Value: &latitude,
		}
		m.Pack = append(m.Pack, *lon)
	}

	return m
}

func (m MessageAccepted) IsLocated() bool {
	return m.Latitude() != 0 && m.Longitude() != 0
}

func (m MessageAccepted) Latitude() float64 {
	for _, r := range m.Pack {
		if r.Unit == senml.UnitLat {
			return *r.Value
		}
	}
	return 0
}

func (m MessageAccepted) Longitude() float64 {
	for _, r := range m.Pack {
		if r.Unit == senml.UnitLon {
			return *r.Value
		}
	}
	return 0
}
