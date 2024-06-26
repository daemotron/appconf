package appconf

import (
	"os"
	"path/filepath"
)

func (conf *AppConf) userDataDir() (string, error) {
	var err error
	var base string
	if conf.Roaming {
		base, err = os.UserConfigDir()
	} else {
		base, err = os.UserCacheDir()
	}
	if err != nil {
		return "", err
	}
	return filepath.Join(base, conf.Author, conf.Name, conf.Version), err
}

func (conf *AppConf) siteDataDir(_ bool) (string, error) {
	base := os.Getenv("ALLUSERSPROFILE")
	if base == "" {
		return "", ErrAllUsersProfileNotDefined
	}
	return filepath.Join(base, conf.Author, conf.Name, conf.Version), nil
}

func (conf *AppConf) globalDataDir() (string, error) {
	return conf.siteDataDir(false)
}

func (conf *AppConf) userConfigDir() (string, error) {
	return conf.userDataDir()
}

func (conf *AppConf) siteConfigDir(_ bool) (string, error) {
	return conf.siteDataDir(false)
}

func (conf *AppConf) globalConfigDir(_ bool) (string, error) {
	return conf.siteConfigDir(false)
}

func (conf *AppConf) userCacheDir() (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, conf.Author, conf.Name, conf.Version, "Cache"), err
}

func (conf *AppConf) globalCacheDir() (string, error) {
	return conf.userCacheDir()
}

func (conf *AppConf) userStateDir() (string, error) {
	return conf.userDataDir()
}

func (conf *AppConf) userLogDir() (string, error) {
	base, err := conf.userDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "Logs"), nil
}
