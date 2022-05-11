package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "note",
		Usage: "Append to latest Apple Notes note",
		ArgsUsage: "<text to append>",
		Description: "Ignores shared notes. Formats as new line by default.",
		HideHelpCommand: true,
		Version: "v0.0.2",
		Flags: []cli.Flag {
		  &cli.BoolFlag{
			Name: "bulleted",
			Aliases: []string{"b"},
			Usage: "Format text as bulleted list item",
		  },
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() < 1 {
				cli.ShowAppHelpAndExit(c, 1)
			}
			text := strings.Join(c.Args().Slice(), " ")
			if c.Bool("bulleted") {
				text = "- " + text
			}
		    fmt.Println(text)
		    return nil
		},
	  }
	
	  err := app.Run(os.Args)
	  if err != nil {
		log.Fatal(err)
	  }
}
