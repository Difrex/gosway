package ipc

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const (
	sway = "sway"
)

// IsSwayAvailable exported wrapper under the checkSway()
func IsSwayAvailable() bool {
	return checkSway() == nil
}

// checkSway checks wayland session and ensure we under the Sway
func checkSway() error {
	swaysock := os.Getenv("SWAYSOCK")
	if swaysock == "" {
		return fmt.Errorf("SWAYSOCK not set")
	}

	err := exec.Command(sway, "--version").Run()
	if err != nil {
		return fmt.Errorf("`sway --version` check failed")
	}
	return nil
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
