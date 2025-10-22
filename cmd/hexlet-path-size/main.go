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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			res, err := pathsize.GetPathSize(cmd.Args().Get(0))
			
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