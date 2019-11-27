package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Database string `yaml:"database"`
}

type CONF struct {
	Production  DatabaseConfig `yaml:"production"`
	Development DatabaseConfig `yaml:"development"`
	Mode        string         `yaml:"mode"`
}

var Config CONF

func init(){
	file, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil{
		panic(err)
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil{
		panic(err)
	}
}