package ipc

import (
	"encoding/json"
)

type Input struct {
	Identifier           string                 `json:"identifier"`
	Name                 string                 `json:"name"`
	Vendor               int64                  `json:"vendor"`
	Product              int64                  `json:"product"`
	Type                 string                 `json:"type"`
	XkbLayoutNames       []string               `json:"xkb_layout_names,omitempty"`
	XkbActiveLayoutIndex int8                   `json:"xkb_active_layout_index"`
	XkbActiveLayoutName  string                 `json:"xkb_active_layout_name"`
	LibInput             map[string]interface{} `json:"libinput"`
}

type inputs []Input

// GetInputs returns all available LibInput devices
func (sc *SwayConnection) GetInputs() ([]Input, error) {
	var inputs []Input

	data, err := sc.SendCommand(IPC_GET_INPUTS, "get_inputs")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &inputs); err != nil {
		return nil, err
	}

	return inputs, nil
}
