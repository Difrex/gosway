package main

import (
	"fmt"
	"os"

	"github.com/Difrex/gosway/ipc"
)

func main() {
	if ctlCommand != "" {
		SendToCTL(ctlCommand)
		os.Exit(0)
	}

	sigWait()

	manager, err := newManager()
	if err != nil {
		panic(err)
	}
	defer manager.store.dbConn.Close()

	go manager.ListenCTL()
	defer cleanUpSocket()

	o, err := manager.listenerConn.SendCommand(ipc.IPC_SUBSCRIBE, "[\"window\"]")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(o))

	ch := make(chan *ipc.Event)
	go manager.listenerConn.SubscribeListener(ch)

	for {
		event := <-ch
		if event.Change == "new" {
			wsConfig, isManaged := manager.isWorkspaceManaged()
			if isManaged {
				fmt.Println(event)
				manager.layouts[wsConfig.Layout].PlaceWindow(event)
			}
		}
	}
}
