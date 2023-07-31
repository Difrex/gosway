package ipc

import (
	"encoding/json"
)

// Workspace represents a workspace
type Workspace struct {
	ID                 int64         `json:"id"`
	Name               string        `json:"name"`
	Rect               Rect          `json:"rect"`
	Focus              []int         `json:"focus"`
	Border             string        `json:"border"`
	CurrentBorderWidth int           `json:"current_border_width"`
	Layout             string        `json:"layout"`
	Orientation        string        `json:"orientation"`
	Percent            interface{}   `json:"percent"`
	WindowRect         WindowRect    `json:"window_rect"`
	DecoRect           DecoRect      `json:"deco_rect"`
	Geometry           Geometry      `json:"geometry"`
	Window             interface{}   `json:"window"`
	Urgent             bool          `json:"urgent"`
	FloatingNodes      []interface{} `json:"floating_nodes"`
	Sticky             bool          `json:"sticky"`
	Num                int           `json:"num"`
	Output             string        `json:"output"`
	Type               string        `json:"type"`
	Representation     string        `json:"representation"`
	Focused            bool          `json:"focused"`
	Visible            bool          `json:"visible"`
}

// GetWorkspaces returns all active workspaces
func (sc *SwayConnection) GetWorkspaces() ([]*Workspace, error) {
	var workspaces []*Workspace

	ws, err := sc.SendCommand(IPC_GET_WORKSPACES, "get_workspaces")
	if err != nil {
		return workspaces, err
	}

	err = json.Unmarshal(ws, &workspaces)
	if err != nil {
		return workspaces, err
	}

	return workspaces, nil
}

// GetFocusedWorkspace returns focused workspace name
func (sc *SwayConnection) GetFocusedWorkspace() (*Workspace, error) {
	var name *Workspace

	workspaces, err := sc.GetWorkspaces()
	if err != nil {
		return name, err
	}

	for i := range workspaces {
		if workspaces[i].Focused {
			name = workspaces[i]
			break
		}
	}

	return name, nil
}
