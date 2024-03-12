package ipc

import (
	"encoding/json"
)

// Tree is a structure that we getting from `get_tree'
type Tree struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	Rect               Rect        `json:"rect"`
	Focused            bool        `json:"focused"`
	Focus              []int       `json:"focus"`
	Border             string      `json:"border"`
	CurrentBorderWidth int         `json:"current_border_width"`
	Layout             string      `json:"layout"`
	Orientation        string      `json:"orientation"`
	Percent            float64     `json:"percent"`
	WindowRect         WindowRect  `json:"window_rect"`
	DecoRect           DecoRect    `json:"deco_rect"`
	Geometry           Geometry    `json:"geometry"`
	Window             interface{} `json:"window"`
	Urgent             bool        `json:"urgent"`
	FloatingNodes      []Node      `json:"floating_nodes"`
	Sticky             bool        `json:"sticky"`
	Type               string      `json:"type"`
	Nodes              []Node      `json:"nodes"`
}

// FindParent returns the parent node of the provided node ID
func (sc *SwayConnection) FindParent(id int64) Node {
	tree, _ := sc.GetTree()
	var result []Node

	parent := finder(result, tree.Nodes, tree.Nodes[0], id)

	return parent[0]
}

// finder recursive finds a top level node
func finder(result, n []Node, p Node, id int64) []Node {
	if len(result) > 0 {
		return result
	}
	for _, node := range n {
		if node.ID == id {
			result = append(result, node)
			return result
		}
		if len(node.Nodes) > 0 {
			finder(result, node.Nodes, node, id)
		}
	}
	return n
}

// FindFocusedNodes finds the all focused nodes and send it to the channel
func FindFocusedNodes(tree []Node, ch chan Node) {
	for _, node := range tree {
		if len(node.Nodes) > 0 {
			FindFocusedNodes(node.Nodes, ch)
		}
		if node.Focused {
			ch <- node
		}
	}
}

// GetTree calls get_tree through a unix socket and return the Tree
func (sc *SwayConnection) GetTree() (*Tree, error) {
	tree := &Tree{}

	data, err := sc.SendCommand(IPC_GET_TREE, "get_tree")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, tree)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type WindowRect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type DecoRect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Geometry struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type WindowProperties struct {
	Class        string      `json:"class"`
	Instance     string      `json:"instance"`
	Title        string      `json:"title"`
	TransientFor interface{} `json:"transient_for"`
	WindowRole   string      `json:"window_role"`
}

type FloatingNodes struct {
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

type Modes struct {
	Width   int `json:"width"`
	Height  int `json:"height"`
	Refresh int `json:"refresh"`
}

type CurrentMode struct {
	Width   int `json:"width"`
	Height  int `json:"height"`
	Refresh int `json:"refresh"`
}

type Node struct {
	ID                 int64            `json:"id"`
	Name               string           `json:"name"`
	Rect               Rect             `json:"rect"`
	Focused            bool             `json:"focused"`
	Focus              []int            `json:"focus"`
	Border             string           `json:"border"`
	CurrentBorderWidth int              `json:"current_border_width"`
	Layout             string           `json:"layout"`
	Orientation        string           `json:"orientation"`
	Percent            float64          `json:"percent"`
	WindowRect         WindowRect       `json:"window_rect"`
	DecoRect           DecoRect         `json:"deco_rect"`
	Geometry           Geometry         `json:"geometry"`
	Window             interface{}      `json:"window"`
	Urgent             bool             `json:"urgent"`
	FloatingNodes      []Node           `json:"floating_nodes"`
	Sticky             bool             `json:"sticky"`
	Type               string           `json:"type"`
	Nodes              []Node           `json:"nodes"`
	Active             bool             `json:"active,omitempty"`
	Primary            bool             `json:"primary,omitempty"`
	Visible            bool             `json:"visible,omitempty"`
	Pid                int              `json:"pid"`
	Make               string           `json:"make,omitempty"`
	Model              string           `json:"model,omitempty"`
	Serial             string           `json:"serial,omitempty"`
	Scale              float64          `json:"scale,omitempty"`
	Transform          string           `json:"transform,omitempty"`
	CurrentWorkspace   string           `json:"current_workspace,omitempty"`
	Modes              []Modes          `json:"modes,omitempty"`
	CurrentMode        CurrentMode      `json:"current_mode,omitempty"`
	Representation     string           `json:"representation,omitempty"`
	WindowProperties   WindowProperties `json:"window_properties"`
}
