package main

import (
	"github.com/cloudfoundry/cli/plugin"

	"github.com/xchapter7x/cf-app-attack/attackplugin"
)

func main() {
	plugin.Start(new(attackplugin.AppAttack))
}
