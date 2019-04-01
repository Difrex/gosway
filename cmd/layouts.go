package main

import (
	"github.com/Difrex/gosway/ipc"
)

type Layout interface {
	PlaceWindow(*ipc.Event) error
	Manage() error
}

func NewLayouts(conn *ipc.SwayConnection, store *store) map[string]Layout {
	layouts := make(map[string]Layout)

	spiral := Layout(NewSpiralLayout(conn, store))
	layouts["spiral"] = spiral

	return layouts
}

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
