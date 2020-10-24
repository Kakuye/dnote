// Package xdg provides an implementation of the XDG Base Directory
// specification depending on the operating system of the user.
package xdg

import (
	"os"
	"os/user"

	"github.com/pkg/errors"
)

// The environment variable names for the XDG base directory specification
var (
	envConfigHome = "XDG_CONFIG_HOME"
	envDataHome   = "XDG_DATA_HOME"
)

var (
	// ConfigHome is the full path to the directory in which user-specific
	// configurations should be written.
	ConfigHome string
	// DataHome is the full path to the directory in which user-specific data
	// files should be written.
	DataHome string
)

func init() {
	homeDir, err := getHomeDir()
	if err != nil {
		panic(errors.Wrap(err, "getting home dir"))
	}

	ConfigHome = xdgPath(envConfigHome, getConfigHome(homeDir))
	DataHome = xdgPath(envDataHome, getDataHome(homeDir))
}

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.Wrap(err, "getting current user")
	}

	return usr.HomeDir, nil
}

func xdgPath(envName, defaultPath string) string {
	if dir := os.Getenv(envName); dir != "" {
		return dir
	}

	return defaultPath
}
