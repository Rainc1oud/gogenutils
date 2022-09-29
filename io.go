package gogenutils

import (
	"io/fs"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// WriteToYaml writes given object to yaml
func WriteToYaml(filePath string, object interface{}, perm fs.FileMode) error {
	bytes, err := yaml.Marshal(object)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filePath, bytes, perm); err != nil {
		return err
	}
	return nil
}
