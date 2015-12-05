package attackplugin

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/xchapter7x/cf-app-attack/vegetaclihelper"
)

var VegetaRunner = vegetaclihelper.VegetaCliExecute

func (c *AppAttack) Run(cliConnection plugin.CliConnection, args []string) {
	switch args[0] {
	case CmdBench:
		vegetaArgs := args[2:]
		VegetaRunner(vegetaArgs)

	default:
		fmt.Println("Invalid command:", args[0])
	}
}

func (c *AppAttack) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: PluginName,
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     CmdBench,
				HelpText: "Run a performance/load test function for an app in cf",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s mycoolappname [vegeta globals] [attack | report | dump] [vegeta args]", CmdBench),
				},
			},
		},
	}
}
