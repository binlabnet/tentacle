package thycotic

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func  (p *Provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.Alias,
		Usage:     "Access secrets and passwords stored in thycotic secret server",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret in thycotic secret server",
				Before: func (ctx *cli.Context) error{
					if !ctx.IsSet("name"){
						return fmt.Errorf("`name` is required argument")
					}
					return nil
				},
				Action: func(c *cli.Context) error {
					p.Authenticate()
					queryData := p.CommandFlagsToQueryData(c)

					return p.Get(queryData)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "secretId",
						Usage:   "Specify the thycotic secret id",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in thycotic secret server",
				Action: func(c *cli.Context) error {
					p.Authenticate()
					queryData := p.CommandFlagsToQueryData(c)

					return p.List(queryData)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "criteria",
						Aliases: []string{"searchTerm"},
						Usage:   "Specify the thycotic search term",
					},
				},
			},
		},
	}
}