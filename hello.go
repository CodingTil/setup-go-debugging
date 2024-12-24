package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	err := (&cli.App{}).Run(os.Args)
	if err != nil {
		return
	}
}
