package dirs

import (
	"path/filepath"

	"github.com/pkg/errors"
)

func init() {
	homeDir := getHomeDir()

	ConfigHome = filepath.Join(homeDir, ".dnote")
	DataHome = filepath.Join(homeDir, ".dnote")
}
