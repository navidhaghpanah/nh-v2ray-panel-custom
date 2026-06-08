// Package global provides global state management for the web panel.
package global

import (
	"github.com/navidhaghpanah/nh-v2ray-panel/v3/sub"
	"github.com/navidhaghpanah/nh-v2ray-panel/v3/web"
)

var (
	webServer *web.Server
	subServer *sub.Server
	restartHook func()
)

// SetWebServer sets the global web server instance.
func SetWebServer(server *web.Server) {
	webServer = server
}

// SetSubServer sets the global sub server instance.
func SetSubServer(server *sub.Server) {
	subServer = server
}

// SetRestartHook sets the restart hook function.
func SetRestartHook(hook func()) {
	restartHook = hook
}
