package main

import (
	"fmt"

	"github.com/Difrex/gosway/ipc"
)

type SpiralLayout struct {
	Conn  *ipc.SwayConnection
	store *store
}

func NewSpiralLayout(conn *ipc.SwayConnection, store *store) *SpiralLayout {
	layout := &SpiralLayout{}
	layout.Conn = conn
	layout.store = store
	return layout
}

func (s *SpiralLayout) PlaceWindow(event *ipc.Event) error {
	nodes, err := s.Conn.GetFocusedWorkspaceWindows()
	if err != nil {
		return err
	}
	var result ipc.Node
	for _, node := range nodes {
		if node.Focused {
			result = node
			break
		}
	}

	if result.WindowRect.Width > result.WindowRect.Height {
		_, err := s.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] split h", event.Container.ID))
		return err
	} else {
		_, err := s.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] split v", event.Container.ID))
		return err
	}

	return nil
}

func (s *SpiralLayout) Manage() error {
	ws, err := s.Conn.GetFocusedWorkspace()
	if err != nil {
		return err
	}

	wc := WorkspaceConfig{
		Name:    ws.Name,
		Layout:  "spiral",
		Managed: true,
	}

	if err := s.store.put([]byte(ws.Name), wc); err != nil {
		return err
	}

	return nil
}
