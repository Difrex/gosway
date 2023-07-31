package ipc

type ChangeEvent string

// Event represents a Sway event
type Event struct {
	Change    ChangeEvent `json:"change"`
	Container Container   `json:"container"`
}

// Container represents the Sway container
type Container struct {
	ID                 int              `json:"id"`
	Name               string           `json:"name"`
	Rect               Rect             `json:"rect"`
	Focused            bool             `json:"focused"`
	Focus              []interface{}    `json:"focus"`
	Border             string           `json:"border"`
	CurrentBorderWidth int              `json:"current_border_width"`
	Layout             string           `json:"layout"`
	Orientation        string           `json:"orientation"`
	Percent            float64          `json:"percent"`
	WindowRect         WindowRect       `json:"window_rect"`
	DecoRect           DecoRect         `json:"deco_rect"`
	Geometry           Geometry         `json:"geometry"`
	Window             int              `json:"window"`
	Urgent             bool             `json:"urgent"`
	FloatingNodes      []FloatingNodes  `json:"floating_nodes"`
	Sticky             bool             `json:"sticky"`
	Type               string           `json:"type"`
	FullscreenMode     int              `json:"fullscreen_mode"`
	Pid                int              `json:"pid"`
	AppID              interface{}      `json:"app_id"`
	Visible            bool             `json:"visible"`
	Marks              []interface{}    `json:"marks"`
	WindowProperties   WindowProperties `json:"window_properties"`
	Nodes              []Node           `json:"nodes"`
}

// GetFocusedWorkspaceWindows returns all the windows from the focused workspace
func (sc *SwayConnection) GetFocusedWorkspaceWindows() ([]Node, error) {
	var nodes []Node

	o, err := sc.GetActiveOutput()
	if err != nil {
		return nodes, err
	}

	tree, err := sc.GetTree()
	if err != nil {
		return nodes, err
	}

	ws, err := sc.GetFocusedWorkspace()
	if err != nil {
		return nodes, err
	}

	for _, node := range tree.Nodes {
		if node.Name == o.Name && node.Active {
			for _, w := range node.Nodes {
				if w.Name == ws.Name {
					nodes = findWindows(w.Nodes)
				}
			}
		}
	}

	return nodes, nil
}

// GetAllFloatingWindows returns all the floating windows including scratchpads
func GetAllFloatingWindows(n []Node) []Node {
	var windows []Node

	for _, node := range n {
		if len(node.FloatingNodes) > 0 {
			windows = append(windows, node.FloatingNodes...)
		} else if len(node.Nodes) > 0 {
			for _, child := range findWindows(node.Nodes) {
				if len(child.FloatingNodes) > 0 {
					windows = append(windows, child.FloatingNodes...)
				}
			}
		}
	}

	return windows
}

// findWindows recursive finding a windows
func findWindows(n []Node) []Node {
	var nodes []Node

	for _, node := range n {
		if node.Name != "" {
			nodes = append(nodes, node)
		} else {
			nodes = append(nodes, findWindows(node.Nodes)...)
		}
	}

	return nodes
}

// GetLargestWindow returns largest window ID
func GetLargestWindowID(n []Node) int64 {
	var id int64
	var max int
	for i := range n {
		if (n[i].Rect.Height + n[i].Rect.Width) > max {
			max = n[i].Rect.Height + n[i].Rect.Width
			id = n[i].ID
		}
	}

	return id
}
