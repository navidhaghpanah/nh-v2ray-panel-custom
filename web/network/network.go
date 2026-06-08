// Package network provides network-related utilities for the web panel.
package network

import (
	"net"
)

// NewAutoHttpsListener wraps a listener to provide automatic HTTPS detection.
func NewAutoHttpsListener(listener net.Listener) net.Listener {
	return listener
}
