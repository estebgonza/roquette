package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/estebgonza/roquette/roquettor"

	"github.com/urfave/cli"
)

const (
	defaultPlanFile     string = "plan.json"
	defaultDatabaseFile string = "database.json"
)

var (
	run = flag.Bool("run", false, "")
)

const usage = `
Usage: roquette [options...]

Options:
  -run  Execute your plan on specified database.
`

func main() {
	app := cli.NewApp()
	app.Name = "Roquette"
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	flag.Parse()
	if !*run {
		usageAndExit("")
	}

	var jsonFile *os.File
	var byteValue []byte

	// Plan
	// TODO: Improve 'json to struct' operations
	var p roquettor.Plan
	jsonFile, err := os.Open(defaultPlanFile)
	if err != nil {
		errAndExit("No plan.json found, please read https://github.com/estebgonza/roquette")
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &p)

	// Database
	// TODO: Improve 'json to struct' operations
	var d roquettor.Database
	jsonFile, err = os.Open(defaultDatabaseFile)
	if err != nil {
		errAndExit("No database.json found, please read https://github.com/estebgonza/roquette")
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &d)

	// Execute plan on specified database
	err = roquettor.Execute(&d, &p)
	if err != nil {
		errAndExit(err.Error())
	}
}

func errAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func usageAndExit(msg string) {
	fmt.Printf(msg)
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
