package main

import (
	"encoding/json"

	"github.com/Difrex/gosway/ipc"
)

type manager struct {
	commandConn  *ipc.SwayConnection
	listenerConn *ipc.SwayConnection
	store        *store
	layouts      map[string]Layout
}

type WorkspaceConfig struct {
	Name    string `json:"name"`
	Layout  string `json:"layout"`
	Managed bool   `json:"managed"`
}

func newManager() (*manager, error) {
	manager := &manager{}

	commandConn, err := ipc.NewSwayConnection()
	if err != nil {
		return manager, err
	}
	manager.commandConn = commandConn

	listenerConn, err := ipc.NewSwayConnection()
	if err != nil {
		return manager, err
	}
	manager.listenerConn = listenerConn

	store, err := newStore()
	if err != nil {
		return manager, err
	}
	manager.store = store

	manager.layouts = NewLayouts(commandConn, store)

	return manager, nil
}

func (m *manager) isWorkspaceManaged() (bool, error) {
	var managed bool
	var config WorkspaceConfig

	ws, err := m.commandConn.GetFocusedWorkspace()
	if err != nil {
		return managed, err
	}

	data, err := m.store.get([]byte(ws.Name))
	if err != nil {
		return managed, nil
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return managed, nil
	}

	return config.Managed, nil
}
