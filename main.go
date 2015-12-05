package main

import (
	"os"

	"github.com/xchapter7x/cf-app-attack/vegetaclihelper"
)

func main() {
	vegetaclihelper.VegetaCliExecute(os.Args[1:])
}
