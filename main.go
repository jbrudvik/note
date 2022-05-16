package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

//go:embed append_to_latest_unshared_note.scpt
var appleScript string

func main() {
	app := &cli.App{
		Name:            "note",
		Usage:           "Append to latest Apple Notes note",
		ArgsUsage:       "\"Text to append\"",
		Description:     "Ignores shared notes. Formats as new line by default.",
		HideHelpCommand: true,
		Version:         "v0.0.7",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "bulleted",
				Aliases: []string{"b"},
				Usage:   "Format text as bulleted list item",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() < 1 {
				cli.ShowAppHelpAndExit(c, 1)
			}

			// Get all text input and format it
			text := strings.Join(c.Args().Slice(), " ")
			if c.Bool("bulleted") {
				text = "b " + text
			}

			// Execute AppleScript
			cmd := exec.Command("osascript", "-e", appleScript, text)
			_, err := cmd.Output()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to run AppleScript to write to note: "+err.Error())
				return nil
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
