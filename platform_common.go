//go:build !windows

package appconf

import (
	"os"
	"os/user"
	"path/filepath"
)

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func (conf *AppConf) userCacheDir() (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, conf.Name, conf.Version), nil
}
