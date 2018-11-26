package defaults

import (
	"log"
	"path/filepath"

	"github.com/gladiusio/gladius-common/pkg/utils"
)

// NetworkDaemonDefaults Returns the network daemon's default config.
func NetworkDaemonDefaults() map[string]string {
	m := make(map[string]string)
	base, err := utils.GetGladiusBase()
	if err != nil {
		log.Fatal(err)
	}

	m["ContentDirectory"] = filepath.Join(base, "content")
	m["ContentPort"] = "8080"
	m["P2PSeedNodeAddress"] = "165.227.16.209"
	m["P2PSeedNodePort"] = "7946"
	m["ControldHostname"] = "localhost"
	m["ControldPort"] = "3001"
	m["ControldProtocol"] = "http"
	m["LogLevel"] = "info"

	return m
}
