package main

import (
	"github.com/Difrex/gosway/ipc"
)

// Layout is an
type Layout interface {
	// PlaceWindow must receive an *ipc.Event
	// and do the container manipulation
	PlaceWindow(*ipc.Event) error
	// Manage must store WorkspaceConfig in the database with
	// the workspace name, layout name and with the Managed: true
	Manage() error
}

// NewLayouts initilizes all the layouts
func NewLayouts(conn *ipc.SwayConnection, store *store) map[string]Layout {
	layouts := make(map[string]Layout)

	spiral := Layout(NewSpiralLayout(conn, store))
	layouts["spiral"] = spiral

	return layouts
}

// Unmanage makes the currently focused workspace non-management
func (m *manager) Unmanage() error {
	ws, err := m.commandConn.GetFocusedWorkspace()
	if err != nil {
		return err
	}

	wc := WorkspaceConfig{
		Name:    ws.Name,
		Layout:  "",
		Managed: false,
	}

	if err := m.store.put([]byte(ws.Name), wc); err != nil {
		return err
	}

	return nil
}

type FiberLayout struct{}
type TopLayout struct{}
type BottomLayout struct{}
type LeftLayout struct{}
type RightLayout struct{}
