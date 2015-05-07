package stackup

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

// Supfile represents the Stackup configuration YAML file.
type Supfile struct {
	Networks map[string]Network  `yaml:"networks"`
	Commands map[string]Command  `yaml:"commands"`
	Targets  map[string][]string `yaml:"targets"`
	Env      map[string]string   `yaml:"env"`
}

// Network is group of hosts with extra custom env vars.
type Network struct {
	Hosts []string          `yaml:"hosts"`
	Env   map[string]string `yaml:"env"`
}

// Command represents command(s) to be run remotely.
type Command struct {
	Name   string   `yaml:-`        // Command name.
	Desc   string   `yaml:"desc"`   // Command description.
	Run    string   `yaml:"run`     // Command(s) to be run remotelly.
	Script string   `yaml:"script"` // Load command(s) from script and run it remotelly.
	Upload []Upload `yaml:"upload"` // See below.
	Stdin  bool     `yaml:"stdin"`  // Attach localhost STDOUT to remote commands' STDIN?
}

// Upload represents file copy operation from localhost Src path to Dst
// path of every host in a given Network.
type Upload struct {
	Src string `yaml:"src"`
	Dst string `yaml:"dst"`
}

// NewSupfile parses configuration file and returns Supfile or error.
func NewSupfile(file string) (*Supfile, error) {
	var conf Supfile
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}