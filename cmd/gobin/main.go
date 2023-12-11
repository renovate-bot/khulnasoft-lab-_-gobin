package main

import (
	"github.com/anchore/clio"
	"github.com/khulnasoft-lab/gobin/cmd/gobin/cli"
)

// applicationName is the non-capitalized name of the application (do not change this)
const (
	applicationName = "gobin"
	notProvided     = "[not provided]"
)

// all variables here are provided as build-time arguments, with clear default values
var (
	version        = notProvided
	buildDate      = notProvided
	gitCommit      = notProvided
	gitDescription = notProvided
)

func main() {
	app := cli.New(
		clio.Identification{
			Name:           applicationName,
			Version:        version,
			BuildDate:      buildDate,
			GitCommit:      gitCommit,
			GitDescription: gitDescription,
		},
	)

	app.Run()
}
