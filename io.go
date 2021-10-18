package gogenutils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// FileExists checks whether filename exists and is a (regular) file (it returns (somehwat peculiar?) true, error if exists but is a dir)
func FileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, err
	}
	if info.IsDir() {
		return true, fmt.Errorf("%s is a directory", filename)
	}
	return true, nil
}

// DirExists checks whether dirname exists and is a dir (it returns (somehwat peculiar?) true, error if exists but is not a dir)
func DirExists(dirname string) (bool, error) {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false, err
	}
	if !info.IsDir() {
		return true, fmt.Errorf("%s is not a directory", dirname)
	}
	return true, nil
}

// WriteToYaml writes given object to yaml
func WriteToYaml(filePath string, object interface{}) error {
	bytes, err := yaml.Marshal(object)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
