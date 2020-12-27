package dirs

import (
	"path/filepath"
)

// The environment variable names for the XDG base directory specification
var (
	envConfigHome = "XDG_CONFIG_HOME"
	envDataHome   = "XDG_DATA_HOME"
)

func init() {
	homeDir := getHomeDir()

	ConfigHome = readPath(envConfigHome, getConfigHome(homeDir))
	DataHome = readPath(envDataHome, getDataHome(homeDir))
}

func getConfigHome(homeDir string) string {
	return filepath.Join(homeDir, ".config")
}

func getDataHome(homeDir string) string {
	return filepath.Join(homeDir, ".local/share")
}
