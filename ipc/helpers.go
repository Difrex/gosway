package ipc

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

const (
	sway = "sway"
)

// checkSway checks wayland session and ensure we under the Sway
func checkSway() bool {
	swaysock := os.Getenv("SWAYSOCK")
	if swaysock != "" && strings.HasSuffix(swaysock, ".sock") {
		return true
	}

	err := exec.Command(sway, "--version").Run()
	if err != nil ||
		os.Getenv("WAYLAND_DISPLAY") == "" ||
		os.Getenv("XDG_SESSION_TYPE") != "wayland" {
		return false
	}
	return true
}

// runSwayCMD runs sway executable with provided args
// returns sway stdout
func runSwayCMD(args ...string) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(sway, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}

	return stdout.String(), nil
}
