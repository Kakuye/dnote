// Package dirs provides base directory definitions for the system
package dirs

import (
	"os"
	"os/user"

	"github.com/pkg/errors"
)

var (
	// ConfigHome is the full path to the directory in which user-specific
	// configurations should be written.
	ConfigHome string
	// DataHome is the full path to the directory in which user-specific data
	// files should be written.
	DataHome string
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(errors.Wrap(err, "getting home dir"))
	}

	return usr.HomeDir
}

func readPath(envName, defaultPath string) string {
	if dir := os.Getenv(envName); dir != "" {
		return dir
	}

	return defaultPath
}
