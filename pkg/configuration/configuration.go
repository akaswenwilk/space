package configuration

import (
	"os"
)

var (
	defaultSpacesDirectory = os.Getenv("HOME") + "/spaces"
	username               = os.Getenv("GIT_USERNAME")
	token                  = os.Getenv("GIT_TOKEN")
)

type Conf struct {
	SpacesDirectory string
	GithubUsername  string
	GithubToken     string
}

func New() Conf {
	return Conf{
		SpacesDirectory: defaultSpacesDirectory,
		GithubUsername:  username,
		GithubToken:     token,
	}
}
