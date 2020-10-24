// +build linux

package xdg

import (
	"path/filepath"
)

func getConfigHome(homeDir string) string {
	return filepath.Join(homeDir, ".config")
}

func getDataHome(homeDir string) string {
	return filepath.Join(homeDir, ".local/share")
}
