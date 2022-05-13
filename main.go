package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

const scriptFile string = "append_to_latest_unshared_note.scpt"

func main() {
	app := &cli.App{
		Name:            "note",
		Usage:           "Append to latest Apple Notes note",
		ArgsUsage:       "<text to append>",
		Description:     "Ignores shared notes. Formats as new line by default.",
		HideHelpCommand: true,
		Version:         "v0.0.5",
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

			// Get all text
			text := strings.Join(c.Args().Slice(), " ")

			// Format text
			if c.Bool("bulleted") {
				text = "<ul><li>" + text + "</li></ul>"
			} else {
				text = "<div>" + text + "</div>"
			}

			// Get AppleScript
			scriptData, err := os.ReadFile(scriptFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not read AppleScript template: "+err.Error())
			}
			scriptString := string(scriptData)

			// Execute AppleScript
			cmd := exec.Command("osascript", "-e", scriptString, text)
			_, err = cmd.Output()
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
