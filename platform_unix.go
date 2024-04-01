//go:build !darwin && !windows

package appconf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (conf *AppConf) userDataDir() (string, error) {
	var base string
	var err error
	base = os.Getenv("XDG_DATA_HOME")
	if base == "" {
		base, err = getHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(base, ".local", "share")
	}
	return filepath.Join(base, conf.Name, conf.Version), nil
}

func (conf *AppConf) siteDataDir(multiPath bool) (string, error) {
	xdg := os.Getenv("XDG_DATA_DIRS")
	if !multiPath {
		if xdg == "" || strings.Contains(xdg, "/usr/local/share") {
			return filepath.Join("/usr/local/share", conf.Name, conf.Version), nil
		}
		if strings.Contains(xdg, "/usr/share") {
			return filepath.Join("/usr/share", conf.Name, conf.Version), nil
		}
		xdgs := strings.Split(xdg, fmt.Sprintf("%c", os.PathListSeparator))
		return filepath.Join(xdgs[0], conf.Name, conf.Version), nil
	}
	if xdg == "" {
		xdg = "/usr/local/share" + fmt.Sprintf("%c", os.PathListSeparator) + "/usr/share"
	}
	xdgs := strings.Split(xdg, fmt.Sprintf("%c", os.PathListSeparator))
	for index, element := range xdgs {
		xdgs[index] = filepath.Join(element, conf.Name, conf.Version)
	}
	return strings.Join(xdgs, fmt.Sprintf("%c", os.PathListSeparator)), nil
}

func (conf *AppConf) globalDataDir() (string, error) {
	return filepath.Join("/var", "lib", conf.Name), nil
}

func (conf *AppConf) userConfigDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, conf.Name, conf.Version), nil
}

func (conf *AppConf) siteConfigDir(multiPath bool) (string, error) {
	xdg := os.Getenv("XDG_CONFIG_DIRS")
	if xdg == "" {
		return filepath.Join("/etc", "xdg", conf.Name, conf.Version), nil
	}
	if !multiPath {
		if strings.Contains(xdg, "/etc/xdg") {
			return filepath.Join("/etc", "xdg", conf.Name, conf.Version), nil
		}
		xdgs := strings.Split(xdg, fmt.Sprintf("%c", os.PathListSeparator))
		return filepath.Join(xdgs[0], conf.Name, conf.Version), nil
	}
	xdgs := strings.Split(xdg, fmt.Sprintf("%c", os.PathListSeparator))
	for index, element := range xdgs {
		xdgs[index] = filepath.Join(element, conf.Name, conf.Version)
	}
	return strings.Join(xdgs, fmt.Sprintf("%c", os.PathListSeparator)), nil
}

func (conf *AppConf) globalConfigDir(multiPath bool) (string, error) {
	if multiPath {
		return strings.Join(
			[]string{filepath.Join("/etc", conf.Name), filepath.Join("/etc", conf.Author)},
			fmt.Sprintf("%c", os.PathListSeparator)), nil
	}
	return filepath.Join("/etc", conf.Name), nil
}

func (conf *AppConf) globalCacheDir() (string, error) {
	return filepath.Join("/var", "cache", conf.Name), nil
}

func (conf *AppConf) userStateDir() (string, error) {
	var base string
	var err error
	base = os.Getenv("XDG_STATE_HOME")
	if base == "" {
		base, err = getHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(base, ".local", "state")
	}
	return filepath.Join(base, conf.Name, conf.Version), nil
}

func (conf *AppConf) userLogDir() (string, error) {
	base, err := conf.userCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "log"), nil
}
