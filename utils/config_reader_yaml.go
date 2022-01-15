package utils

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type yamlConfigReader struct {
	Dir  string
	File string
	buf  *bytes.Buffer
}

// NewYamlConfigReader create yaml config reader instance with file directory
// and the file path.
func NewYamlConfigReader(dir, file string) ConfigReader {
	return &yamlConfigReader{
		Dir:  dir,
		File: file,
	}
}

func (y *yamlConfigReader) Read(i interface{}) (err error) {
	f := make([]byte, 0)
	f, err = ioutil.ReadFile(fmt.Sprintf("%v/%v", y.Dir, y.File))
	if err != nil {
		return
	}
	err = yaml.Unmarshal(f, i)
	if err != nil {
		return
	}
	return
}
