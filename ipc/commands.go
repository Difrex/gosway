package ipc

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	swayMsg = "swaymsg"
)

// RunSwayCommand returns STDOUT or STDERR of the swaymsg
func (sc *SwayConnection) RunSwayCommand(cmd string) ([]byte, error) {
	return swaymsg(cmd)
}

// swaymsg runs swaymsg with the provided message string
// returns STDOUT if the swaymsg process exits with the 0 exit code else return a STDERR
func swaymsg(msg string) ([]byte, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(swayMsg, msg)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.Bytes(), err
	}

	return stdout.Bytes(), nil
}

// Command runs an sway command with the node ID: [con_id]=n.ID com
func (n Node) Command(com string) ([]byte, error) {
	return swaymsg(fmt.Sprintf("[con_id=%d] %s", n.ID, com))
}
