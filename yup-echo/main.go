package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/echo"
)

const (
	flagNoNewline = "no-newline"
	flagEscape    = "escape"
)

func main() {
	app := &cli.App{
		Name:  "echo",
		Usage: "display a line of text",
		UsageText: `echo [OPTIONS] [STRING...]

   Echo the STRING(s) to standard output.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagNoNewline,
				Aliases: []string{"n"},
				Usage:   "do not output the trailing newline",
			},
			&cli.BoolFlag{
				Name:    flagEscape,
				Aliases: []string{"e"},
				Usage:   "enable interpretation of backslash escapes",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments as strings
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.Bool(flagNoNewline) {
		params = append(params, NoNewline)
	}
	if c.Bool(flagEscape) {
		params = append(params, EnableEscape)
	}

	// Create and execute the echo command
	cmd := Echo(params...)
	return gloo.Run(cmd)
}
