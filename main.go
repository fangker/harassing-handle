package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"time"
)

type AppParamConfig struct {
	phone string
	name  string
}

func main() {
	app := &cli.App{
		Name:     "骚扰电话提交小助手",
		Version:  "v1.0.0",
		Usage:    "骚扰电话",
		Compiled: time.Now(),
		Before: func(c *cli.Context) error {
			_, _ = fmt.Fprintf(c.App.Writer, "请勿用于非法用途,骚扰电话将提交到其他具有骚扰性质的电话后台。\n")
			return nil
		},
	}
	apc := &AppParamConfig{}
	app.Commands = []*cli.Command{
		{
			Name:  "submit",
			Usage: "submit a request on harassing webs",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Destination: &apc.name, Value: ""},
				&cli.StringFlag{Name: "phone", Aliases: []string{"p"}, Destination: &apc.phone, Value: ""},
			},
		}}
	app.Action = func(c *cli.Context) error {
		fmt.Println("loading configs...")
		fmt.Println("name:", c.String("name"))
		fmt.Println("phone", c.String("phone"))
		// TODO: http proxy

		return nil
	}
	app.Before = func(c *cli.Context) error {
		fmt.Println("app Before")
		return nil
	}
	app.After = func(c *cli.Context) error {
		fmt.Println("app After")
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
