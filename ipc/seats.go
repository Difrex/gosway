package ipc

import (
	"encoding/json"
)

type Seat struct {
	Name         string  `json:"name"`
	Capabilities int64   `json:"capabilities"`
	Focus        int64   `json:"focus"`
	Devices      []Input `json:"devices"`
}

// GetSeats returns all available seats
func (sc *SwayConnection) GetSeats() ([]Seat, error) {
	var seats []Seat

	data, err := sc.SendCommand(IPC_GET_SEATS, "get_seats")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &seats); err != nil {
		return nil, err
	}

	return seats, nil
}
