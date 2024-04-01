package appconf

import (
	"os"
	"path/filepath"
	"strings"
)

// isFile is an auxiliary function to check whether a path represents a file
func isFile(path string) bool {
	file, err := filepath.EvalSymlinks(path)
	if err != nil {
		return false
	}
	stat, err := os.Stat(file)
	if err == nil {
		return !stat.IsDir()
	}
	return false
}

// ConfigFiles returns a list of all detected configuration files for this application
func (conf *AppConf) ConfigFiles() ([]string, error) {
	var result []string
	files := []string{"config.json", "conf.json", strings.ToLower(conf.Name) + ".json"}
	dirs, err := conf.ConfigDirs(true)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		for _, file := range files {
			candidate := filepath.Join(dir, file)
			if isFile(candidate) {
				result = append(result, candidate)
			}
		}
	}
	return result, nil
}

// UpdateFromFiles updates configuration options from all detected configuration files.
//
// WARNING:
//
//	This method does not guarantee a certain order. If two configuration files contain
//	the same keys, it is random which of the files is parsed last, overwriting values
//	read from files parsed earlier.
func (conf *AppConf) UpdateFromFiles() error {
	cfgFiles, err := conf.ConfigFiles()
	if err != nil {
		return err
	}
	for _, file := range cfgFiles {
		err = conf.updateFromJsonFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}
