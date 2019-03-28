package main

import (
	"fmt"
	"os/exec"

	"github.com/Difrex/gosway/ipc"
)

func main() {

	manager, err := newManager()
	if err != nil {
		panic(err)
	}

	defer manager.store.dbConn.Close()

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
			manager.layouts["spiral"].PlaceWindow(event)
		}
	}
	// tree, _ := sc.GetTree()

	// ch := make(chan ipc.Node)
	// go ipc.FindFocusedNodes(tree.Nodes, ch)

	// result := <-ch

	// if result.WindowRect.Width > result.WindowRect.Height {
	// 	swaymsg(fmt.Sprintf("[con_id=%d] split h", result.ID))
	// 	if len(os.Args) > 1 {
	// 		fmt.Println(os.Args)
	// 		execCMD(os.Args[1], os.Args[1:]...)
	// 	}
	// } else {
	// 	swaymsg(fmt.Sprintf("[con_id=%d] split v", result.ID))
	// 	if len(os.Args) > 1 {
	// 		fmt.Println(os.Args)
	// 		execCMD(os.Args[1], os.Args[1:]...)
	// 	}
	// }
}

func execCMD(c string, args ...string) {
	cmd := exec.Command(c, args...)
	cmd.Run()
}

func swaymsg(msg string) {
	cmd := exec.Command("swaymsg", msg)
	cmd.Run()
}
