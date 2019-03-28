package main

import (
	"fmt"
	"os"

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
	tree, err := s.Conn.GetTree()
	if err != nil {
		return err
	}

	ch := make(chan ipc.Node)
	go ipc.FindFocusedNodes(tree.Nodes, ch)

	result := <-ch

	if result.WindowRect.Width > result.WindowRect.Height {
		swaymsg(fmt.Sprintf("[con_id=%d] split h", result.ID))
		if len(os.Args) > 1 {
			fmt.Println(os.Args)
			// execCMD(os.Args[1], os.Args[1:]...)
		}
	} else {
		swaymsg(fmt.Sprintf("[con_id=%d] split v", result.ID))
		if len(os.Args) > 1 {
			fmt.Println(os.Args)
			// execCMD(os.Args[1], os.Args[1:]...)
		}
	}

	return nil
}

func (s *SpiralLayout) Change() error {
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
