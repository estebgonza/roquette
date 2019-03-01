package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/estebgonza/roquette/roquettor"

	"github.com/urfave/cli"
)

const (
	appName        string = "Roquette"
	appDescription string = "Stress your database"
	appVersion     string = "0.1"

	defaultPlanFile     string = "plan.json"
	defaultDatabaseFile string = "database.json"
)

const helpTemplate = `
Usage: {{.HelpName}} [command]

{{if .Commands}}Commands:

{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}
`

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(helpTemplate)
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appDescription
	app.Version = appVersion

	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Execute your plan on specified database",
			Action: run,
		},
		{
			Name:  "init",
			Usage: "Initialize Roquette files template",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	var jsonFile *os.File
	var byteValue []byte

	// Plan
	// TODO: Improve 'json to struct' operations
	var p roquettor.Plan
	jsonFile, err := os.Open(defaultPlanFile)
	if err != nil {
		return errors.New("No plan.json found, please read https://github.com/estebgonza/roquette")
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &p)

	// Database
	// TODO: Improve 'json to struct' operations
	var d roquettor.Database
	jsonFile, err = os.Open(defaultDatabaseFile)
	if err != nil {
		return errors.New("No database.json found, please read https://github.com/estebgonza/roquette")
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &d)

	// Execute plan on specified database
	return roquettor.Execute(&d, &p)
}
