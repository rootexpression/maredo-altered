package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
)

func initCli() cli.App {
	cwd, _ := os.Getwd()
	cwd = filepath.Base(cwd)

	app := cli.App{
		Name:      "maredo",
		Usage:     "render a markdown file to static HTML",
		Version:   VERSION,
		Copyright: "©️ 2024 dogue <https://github.com/dogue>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "render `FILE` to HTML",
				Value:       "README.md",
				Destination: &sourceFile,
			},
			&cli.StringFlag{
				Name:        "title",
				Aliases:     []string{"t"},
				Usage:       "set output page title to `TITLE`",
				Value:       cwd,
				Destination: &data.Title,
			},
			&cli.StringSliceFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Usage:   "set highlighting languages to be included (see Highlight.js for available languages)",
				Action: func(ctx *cli.Context, s []string) error {
					data.Langs = s
					return nil
				},
			},
			&cli.StringFlag{
				Name:        "page-theme",
				Usage:       "set CSS theme (see https://github.com/dogue/maredo/themes)",
				Value:       "default",
				Destination: &data.PageTheme,
			},
			&cli.StringFlag{
				Name:        "syntax-theme",
				Usage:       "set syntax highlighting theme (see Highlight.js for available themes)",
				Value:       "github-dark",
				Destination: &data.SyntaxTheme,
			},
		},
		Action: func(ctx *cli.Context) error { return nil },
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	return app
}
