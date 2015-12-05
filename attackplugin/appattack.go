package attackplugin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/xchapter7x/cf-app-attack/vegetaclihelper"
)

var VegetaRunner = vegetaclihelper.VegetaCliExecute

func (c *AppAttack) Run(cliConnection plugin.CliConnection, args []string) {
	switch args[0] {
	case CmdBench:
		vegetaArgs := args[2:]
		appHost := "localhost"
		VegetaRunner(vegetaArgs, appHost)

	default:
		fmt.Println("Invalid command:", args[0])
	}
}

func (c *AppAttack) GetVersionType() plugin.VersionType {
	versionArray := strings.Split(strings.TrimPrefix(c.Version, "v"), ".")
	major, _ := strconv.Atoi(versionArray[0])
	minor, _ := strconv.Atoi(versionArray[1])
	build, _ := strconv.Atoi(versionArray[2])
	return plugin.VersionType{
		Major: major,
		Minor: minor,
		Build: build,
	}
}

func (c *AppAttack) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:    PluginName,
		Version: c.GetVersionType(),
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
