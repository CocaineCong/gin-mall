package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// 这个文件是为了方便写的test文件来读取config

type IReader interface {
	readConfig() ([]byte, error)
}

type ConfigReader struct {
	FileName string
}

// 'reader' implementing the Interface
// Function to read from actual file
func (r *ConfigReader) readConfig() ([]byte, error) {
	file, err := ioutil.ReadFile(r.FileName)

	if err != nil {
		log.Fatal(err)
	}
	return file, err
}

func InitConfigForTest(reader IReader) {
	file, err := reader.readConfig()
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		panic(err)
	}
}
