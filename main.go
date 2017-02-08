package main

import (
	"fmt"
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v2"

	"github.com/jroimartin/gocui"
)

func main() {
	app := &cli.App{
		Name:    "rewind",
		Usage:   `a kafka player like tape control`,
		Version: "1.0",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "brokers, b",
				Value: cli.NewStringSlice("localhost:9092"),
				Usage: "kafka brokers address",
			},
		},
		Action: processor,
	}
	app.Run(os.Args)
}

func processor(c *cli.Context) error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("topics", 0, 0, 10, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Topics"
		fmt.Fprintln(v, "TODO:list topic")
	}

	if v, err := g.SetView("control", 11, maxY-10, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Control Panel"
		fmt.Fprintln(v, "TODO: rewind/replay/fastforward")
	}

	if v, err := g.SetView("view", 11, 0, maxX-1, maxY-11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "View"
		fmt.Fprintln(v, "TODO: Display kafka content ")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
