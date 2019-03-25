package ipc

type ChangeEvent string

type Event struct {
	Change    ChangeEvent `json:"change"`
	Container Container   `json:"container"`
}

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

func findWindows(n []Node) []Node {
	var nodes []Node

	for _, node := range n {
		if node.Name != "" {
			nodes = append(nodes, node)
		} else {
			for _, child := range findWindows(node.Nodes) {
				nodes = append(nodes, child)
			}
		}
	}

	return nodes
}
