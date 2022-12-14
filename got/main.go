package main

import (
	"log"
	"os"

	"github.com/harrisonpim/got/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   "got",
		Usage:                  "git in go",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Create an empty Got repository or reinitialize an existing one",
				Action: func(cCtx *cli.Context) error {
					path := cCtx.Args().First()
					cmd.InitNewRepo(path)
					return nil
				},
			},
			{
				Name:  "commit",
				Usage: "Record changes to the repository",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
				},
				Action: func(cCtx *cli.Context) error {
					path := cCtx.Args().Get(0)
					message := cCtx.String("message")
					if err := cmd.Commit(path, message); err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
