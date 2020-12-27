package dirs

import (
	"os"
	"path/filepath"
)

func init() {
	homeDir := getHomeDir()

	ConfigHome = filepath.Join(homeDir, ".dnote")
	DataHome = filepath.Join(homeDir, ".dnote")
}
