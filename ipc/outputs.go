package ipc

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
