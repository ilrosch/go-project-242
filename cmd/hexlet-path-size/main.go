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
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "recursive",
				Usage: "recursive size of directories",
				Aliases: []string{"r"},
				DefaultText: "false",
			},
            &cli.BoolFlag{
                Name:  "human",
                Usage: "human-readable sizes (auto-select unit)",
				Aliases: []string{"H"},
				DefaultText: "false",
            },
			&cli.BoolFlag{
                Name:  "all",
                Usage: "include hidden files and directories",
				Aliases: []string{"a"},
				DefaultText: "false",
            },
        },
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			recursive := cmd.Bool("recursive")
			humanReadable := cmd.Bool("human")
			hiddenFiles := cmd.Bool("all")

			res, err := pathsize.GetPathSize(path, recursive, humanReadable, hiddenFiles)
			
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