package ipc

import (
	"encoding/json"
)

// Output represents structure that we getting from `get_outputs'
type Output struct {
	ID                 int           `json:"id,omitempty"`
	Name               string        `json:"name"`
	Rect               Rect          `json:"rect"`
	Focus              []int         `json:"focus,omitempty"`
	Border             string        `json:"border,omitempty"`
	CurrentBorderWidth int           `json:"current_border_width,omitempty"`
	Layout             string        `json:"layout,omitempty"`
	Orientation        string        `json:"orientation,omitempty"`
	Percent            float64       `json:"percent"`
	WindowRect         WindowRect    `json:"window_rect,omitempty"`
	DecoRect           DecoRect      `json:"deco_rect,omitempty"`
	Geometry           Geometry      `json:"geometry,omitempty"`
	Window             interface{}   `json:"window,omitempty"`
	Urgent             bool          `json:"urgent,omitempty"`
	FloatingNodes      []interface{} `json:"floating_nodes,omitempty"`
	Sticky             bool          `json:"sticky,omitempty"`
	Type               string        `json:"type"`
	Active             bool          `json:"active"`
	Primary            bool          `json:"primary"`
	Make               string        `json:"make"`
	Model              string        `json:"model"`
	Serial             string        `json:"serial"`
	Scale              float64       `json:"scale,omitempty"`
	Transform          string        `json:"transform,omitempty"`
	CurrentWorkspace   string        `json:"current_workspace"`
	Modes              []Modes       `json:"modes"`
	CurrentMode        CurrentMode   `json:"current_mode,omitempty"`
	Focused            bool          `json:"focused,omitempty"`
}

// GetActiveOutput returns the currently active and focused output
func (sc *SwayConnection) GetActiveOutput() (*Output, error) {
	var output *Output
	var o []*Output
	os, err := sc.SendCommand(IPC_GET_OUTPUTS, "get_ouputs")
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(os, &o)
	if err != nil {
		return output, err
	}

	for i, _ := range o {
		if o[i].Focused {
			output = o[i]
			break
		}
	}

	return output, nil
}
