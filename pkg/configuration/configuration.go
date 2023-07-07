package configuration

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	defaultSpacesDirectory = os.Getenv("HOME") + "/spaces"
	spacesYAMLLocation     = os.Getenv("HOME") + "/.spaces.yml"
	username               = os.Getenv("GIT_USERNAME")
	token                  = os.Getenv("GIT_TOKEN")
)

type Conf struct {
	SpacesDirectory string
	Spaces          []string
	GithubUsername  string
	GithubToken     string
}

func New() (Conf, error) {
	spaces, err := LoadSpacesYAML()
	if err != nil {
		return Conf{}, err
	}

	return Conf{
		SpacesDirectory: defaultSpacesDirectory,
		Spaces:          spaces,
		GithubUsername:  username,
		GithubToken:     token,
	}, nil
}

type Spaces struct {
	Repos []string `yaml:"repos"`
}

func LoadSpacesYAML() ([]string, error) {
	file, err := os.ReadFile(spacesYAMLLocation)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml file: %w", err)
	}

	s := Spaces{}

	err = yaml.Unmarshal(file, &s)

	return s.Repos, err
}
