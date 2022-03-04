package configutils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func UnmarshalYamlFile(file string, ptr interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, ptr)
}
