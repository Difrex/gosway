package ipc

import (
	"bytes"
	"os/exec"
)

const (
	swayMsg = "swaymsg"
)

func (sc *SwayConnection) RunSwayCommand(cmd string) ([]byte, error) {
	return swaymsg(cmd)
}

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
