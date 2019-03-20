package ipc

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"unsafe"
)

type SWAY_IPC_GET_WORKSPACES int
type SWAY_IPC_SUBSCRIBE int
type SWAY_IPC_GET_OUTPUTS int
type SWAY_IPC_GET_TREE int
type SWAY_IPC_GET_MARKS int
type SWAY_IPC_GET_BAR_CONFIG int
type SWAY_IPC_GET_VERSION int
type SWAY_IPC_GET_BINDING_MODES int
type SWAY_IPC_GET_CONFIG int
type SWAY_IPC_SEND_TICK int
type SWAY_IPC_SYNC int

// raw and readSwayResponse inspired in github.com/mdirkse/i3ipc-go

const (
	// Magic string for the IPC API.
	MAGICK string = "i3-ipc"
	// The length of the i3 message "header" is 14 bytes: 6 for the _Magic
	// string, 4 for the length of the payload (int32 in native byte order) and
	// another 4 for the message type (also int32 in NBO).
	HEADERLEN = 14
)

type SwayConnection struct {
	Conn net.Conn
	SWAY_IPC_GET_WORKSPACES
	SWAY_IPC_SUBSCRIBE
	SWAY_IPC_GET_OUTPUTS
	SWAY_IPC_GET_TREE
	SWAY_IPC_GET_MARKS
	SWAY_IPC_GET_BAR_CONFIG
	SWAY_IPC_GET_VERSION
	SWAY_IPC_GET_BINDING_MODES
	SWAY_IPC_GET_CONFIG
	SWAY_IPC_SEND_TICK
	SWAY_IPC_SYNC
}

// SendCommand sends command to the Sway unix socket
func (sc *SwayConnection) SendCommand(command int) ([]byte, error) {
	return sc.raw(command, "get_tree")
}

func (sc *SwayConnection) raw(messageType int, args string) ([]byte, error) {
	// Set up the parts of the message.
	var (
		message  = []byte(MAGICK)
		payload  = []byte(args)
		length   = int32(len(payload))
		bytelen  [4]byte
		bytetype [4]byte
	)

	// Black Magic™.
	bytelen = *(*[4]byte)(unsafe.Pointer(&length))
	bytetype = *(*[4]byte)(unsafe.Pointer(&messageType))

	for _, b := range bytelen {
		message = append(message, b)
	}
	for _, b := range bytetype {
		message = append(message, b)
	}
	for _, b := range payload {
		message = append(message, b)
	}

	_, err := sc.Conn.Write(message)
	if err != nil {
		return []byte(nil), err
	}

	msg, err := sc.readSwayResponse()
	if err != nil {
		return []byte(nil), err
	}
	return msg, nil
}

func (sc *SwayConnection) readSwayResponse() ([]byte, error) {
	header := make([]byte, HEADERLEN)
	n, err := sc.Conn.Read(header)

	// Check if this is a valid i3 message.
	if n != HEADERLEN || err != nil {
		return []byte(nil), err
	}

	magicString := string(header[:len(MAGICK)])
	if magicString != MAGICK {
		err = fmt.Errorf(
			"Invalid magic string: got %q, expected %q.",
			magicString, MAGICK)
		return []byte(nil), err
	}

	var bytelen [4]byte
	// Copy the byte values from the slice into the byte array. This is
	// necessary because the address of a slice does not point to the actual
	// values in memory.
	for i, b := range header[len(MAGICK) : len(MAGICK)+4] {
		bytelen[i] = b
	}
	length := *(*int32)(unsafe.Pointer(&bytelen))

	payload := make([]byte, length)
	n, err = sc.Conn.Read(payload)
	if n != int(length) || err != nil {
		return []byte(nil), err
	}

	// Figure out the type of message.
	var bytetype [4]byte
	for i, b := range header[len(MAGICK)+4 : len(MAGICK)+8] {
		bytetype[i] = b
	}

	return payload, err
}

// NewSwayConnection initializes an new Sway connection through unix socket
func NewSwayConnection() (*SwayConnection, error) {
	swayConn := &SwayConnection{}
	path, err := GetSocketPath()
	if err != nil {
		return swayConn, err
	}

	conn, err := net.Dial("unix", path)
	if err != nil {
		return swayConn, err
	}

	swayConn.SWAY_IPC_GET_WORKSPACES = 1
	swayConn.SWAY_IPC_SUBSCRIBE = 2
	swayConn.SWAY_IPC_GET_OUTPUTS = 3
	swayConn.SWAY_IPC_GET_TREE = 4
	swayConn.SWAY_IPC_GET_MARKS = 5
	swayConn.SWAY_IPC_GET_BAR_CONFIG = 6
	swayConn.SWAY_IPC_GET_VERSION = 7
	swayConn.SWAY_IPC_GET_BINDING_MODES = 8
	swayConn.SWAY_IPC_GET_CONFIG = 9
	swayConn.SWAY_IPC_SEND_TICK = 10
	swayConn.SWAY_IPC_SYNC = 11

	swayConn.Conn = conn
	return swayConn, nil
}

// GetSocketPath returns socket path of the running Sway
func GetSocketPath() (string, error) {
	var path string
	if !checkSway() {
		return path, errors.New("Not under the wayland or the Sway executable not found")
	}
	path, err := runSwayCMD("--get-socketpath")
	if err != nil {
		return "", err
	}

	return strings.TrimRight(path, "\n"), nil
}

// ConnectToSocket connects to the Wayland socket
func ConnectToSocket() (net.Conn, error) {
	var conn net.Conn
	path, err := GetSocketPath()
	if err != nil {
		return conn, err
	}

	return net.Dial("unix", path)
}
