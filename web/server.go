// Package web provides web server functionality for the nh-v2ray-panel.
package web

import (
	"context"
	"embed"
	"io/fs"
)

//go:embed dist
var distFS embed.FS

// Server represents the main web server.
type Server struct{}

// NewServer creates a new web server instance.
func NewServer() *Server {
	return &Server{}
}

// Start starts the web server.
func (s *Server) Start() error {
	return nil
}

// Stop stops the web server.
func (s *Server) Stop() error {
	return nil
}

// StopPanelOnly stops only the panel part of the web server.
func (s *Server) StopPanelOnly() error {
	return nil
}

// StartPanelOnly starts only the panel part of the web server.
func (s *Server) StartPanelOnly() error {
	return nil
}

// RestartXray restarts the xray core.
func (s *Server) RestartXray() error {
	return nil
}

// EmbeddedDist returns the embedded dist filesystem.
func EmbeddedDist() fs.FS {
	subFS, _ := fs.Sub(distFS, "dist")
	return subFS
}
