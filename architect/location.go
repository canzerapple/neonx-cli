package architect

import (
	"fmt"
	"os"
	. "path/filepath"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type Location string

var (
	emptyLocation = Location("")
)

//convert a path to location with abs path
func ToLocation(root string) (location Location, err error) {

	root, err = Abs(root)

	if err != nil {
		return emptyLocation, fmt.Errorf("convert [%s] to location fail:%s ", root, err)
	}

	return Location(root), nil
}

func (m Location) Child(location string) Location {

	var (
		names = []string{string(m)}
	)

	names = append(names, strings.Split(location, "/")...)

	return Location(Join(names...))
}

func (m Location) String() string {
	return string(m)
}

func (m Location) GetInfo() (os.FileInfo, bool, error) {

	info, err := os.Stat(string(m))

	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return info, true, err
}

func (m Location) IsExists() (exists bool, err error) {
	_, exists, err = m.GetInfo()
	return
}

func (m Location) IsDirectory() (bool, error) {
	info, exists, err := m.GetInfo()

	if err != nil {
		return false, err
	}

	return exists && info.IsDir(), nil
}

func (m Location) Open(flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(string(m), flag, perm)
}

func (m Location) ReadYAML(v interface{}) error {

	file, err := m.Open(os.O_RDONLY, 0755)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(v)

	if err != nil {
		return fmt.Errorf(
			"decode yaml file [%s] to <%s> fail: %s ",
			m,
			reflect.TypeOf(v),
			err)
	}

	return nil
}

func (m Location) SaveYAML(v interface{}) error {

	file, err := m.Open(os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := yaml.NewEncoder(file)

	err = encoder.Encode(v)

	if err != nil {
		return fmt.Errorf(
			"encode yaml file [%s] to <%s> fail: %s ",
			m,
			reflect.TypeOf(v),
			err)
	}

	return nil
}
