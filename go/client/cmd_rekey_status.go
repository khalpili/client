// Copyright 2016 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
)

type CmdRekeyStatus struct {
	libkb.Contextified
}

func NewCmdRekeyStatus(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:  "status",
		Usage: "Get pending rekey status",
		Action: func(c *cli.Context) {
			cmd := &CmdRekeyStatus{Contextified: libkb.NewContextified(g)}
			cl.ChooseCommand(cmd, "status", c)
		},
	}
}

func (c *CmdRekeyStatus) ParseArgv(ctx *cli.Context) error {
	if len(ctx.Args()) > 0 {
		return UnexpectedArgsError("status")
	}
	return nil
}

func (c *CmdRekeyStatus) Run() error {
	return nil
}

func (c *CmdRekeyStatus) GetUsage() libkb.Usage {
	return libkb.Usage{
		API:    true,
		Config: true,
	}
}
