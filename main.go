package main

import (
	"github.com/cloudfoundry/cli/plugin"

	"github.com/xchapter7x/civet-coffee/attackplugin"
)

var (
	Version string
)

func main() {
	appAttack := &attackplugin.AppAttack{Version: Version}
	plugin.Start(appAttack)
}
