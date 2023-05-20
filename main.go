package main

import (
	_ "embed"

	"github.com/SantaHub/polygon-edge/command/root"
	"github.com/SantaHub/polygon-edge/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
