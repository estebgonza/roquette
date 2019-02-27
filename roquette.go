package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/estebgonza/Roquette/roquettor"

	"github.com/urfave/cli"
)

const (
	defaultPlanFile     string = "plan.json"
	defaultDatabaseFile string = "database.json"
)

var (
	host     = flag.String("m", "GET", "")
	port     = flag.String("h", "", "")
	user     = flag.String("d", "", "")
	password = flag.String("w", "", "")

	c = flag.Int("c", 50, "")
	n = flag.Int("n", 1, "")
)

const usage = `
Usage: roquet [options...] <sql>

Options:
  -h  Hostname of Hive server. Default is 'localhost'.
  -p  Port of Hive server. Default is 10000.
  -n  Number of queries to run. Default is 1.
`

func main() {
	app := cli.NewApp()
	app.Name = "Roquette"
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}
	num := *n
	conc := *c
	if num <= 0 || conc <= 0 {
		usageAndExit("-n and -c cannot be smaller than 1.")
	}

	var jsonFile *os.File
	var byteValue []byte

	// Plan
	var p roquettor.Plan
	jsonFile, _ = os.Open(defaultPlanFile)
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &p)

	// Database
	var d roquettor.Database
	jsonFile, _ = os.Open(defaultDatabaseFile)
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &d)

	// Execute plan on specified database
	roquettor.Execute(&d, &p)
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

type headerSlice []string

func (h *headerSlice) String() string {
	return fmt.Sprintf("%s", *h)
}

func (h *headerSlice) Set(value string) error {
	*h = append(*h, value)
	return nil
}
