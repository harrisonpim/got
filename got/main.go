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
				Usage: "Create an empty Git repository or reinitialize an existing one",
				Action: func(cCtx *cli.Context) error {
					path := cCtx.Args().First()
					cmd.InitNewRepo(path)
					return nil
				},
			},
			{
				Name:  "commit",
				Usage: "Record changes to the repository",
				Action: func(cCtx *cli.Context) error {
					path := cCtx.Args().First()
					cmd.Commit(path)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
