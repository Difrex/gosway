package main

import (
	"github.com/Difrex/gosway/ipc"
)

type Layout interface {
	PlaceWindow(*ipc.Event) error
	Change() error
}

func NewLayouts(conn *ipc.SwayConnection, store *store) map[string]Layout {
	layouts := make(map[string]Layout)

	spiral := Layout(NewSpiralLayout(conn, store))
	layouts["spiral"] = spiral

	return layouts
}

type FiberLayout struct{}
type TopLayout struct{}
type BottomLayout struct{}
type LeftLayout struct{}
type RightLayout struct{}
