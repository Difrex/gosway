package main

import (
	"fmt"
	"os/exec"

	"os"

	"github.com/Difrex/gosway/ipc"
)

func main() {
	sc, err := ipc.NewSwayConnection()
	if err != nil {
		panic(err)
	}

	tree, _ := sc.GetTree()

	ch := make(chan ipc.Node)
	go ipc.FindFocusedNodes(tree.Nodes, ch)

	result := <-ch

	if result.WindowRect.Width > result.WindowRect.Height {
		swaymsg(fmt.Sprintf("[con_id=%d] split h", result.ID))
		if len(os.Args) > 1 {
			fmt.Println(os.Args)
			execCMD(os.Args[1], os.Args[1:]...)
		}
	} else {
		swaymsg(fmt.Sprintf("[con_id=%d] split v", result.ID))
		if len(os.Args) > 1 {
			fmt.Println(os.Args)
			execCMD(os.Args[1], os.Args[1:]...)
		}
	}
}

func execCMD(c string, args ...string) {
	cmd := exec.Command(c, args...)
	cmd.Run()
}

func swaymsg(msg string) {
	cmd := exec.Command("swaymsg", msg)
	cmd.Run()
}
