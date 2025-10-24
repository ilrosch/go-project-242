package main

import (
	pathsize "code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "hexlet-path-size",
		Usage: "print size of a file or directory;",
		Flags: []cli.Flag{
            &cli.BoolFlag{
                Name:  "human",
                Usage: "human-readable sizes (auto-select unit)",
				Aliases: []string{"H"},
				DefaultText: "false",
            },
        },
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			humanReadable := cmd.Bool("human")

			res, err := pathsize.GetPathSize(path, humanReadable)
			
			if err != nil {
				return err
			}
           
			fmt.Println(res)
            return nil
        },
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}