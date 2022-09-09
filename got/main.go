package main

import (
	"log"
	"os"

	"github.com/harrisonpim/got/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "got",
		Usage: "git, in go",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialise an empty git repository in the current directory",
				Action: func(cCtx *cli.Context) error {
					path := &cCtx.Args().First(), "."
					cmd.InitNewRepo(path)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
