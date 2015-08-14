// Copyright (C) 2015 Scaleway. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package cli

import "github.com/scaleway/scaleway-cli/pkg/commands"

var cmdRestart = &Command{
	Exec:        runRestart,
	UsageLine:   "restart [OPTIONS] SERVER [SERVER...]",
	Description: "Restart a running server",
	Help:        "Restart a running server.",
}

func init() {
	cmdRestart.Flag.BoolVar(&restartHelp, []string{"h", "-help"}, false, "Print usage")
}

// Flags
var restartHelp bool // -h, --help flag

func runRestart(cmd *Command, rawArgs []string) error {
	if restartHelp {
		cmd.PrintUsage()
	}
	if len(rawArgs) < 1 {
		cmd.PrintShortUsage()
	}

	args := commands.RestartArgs{
		Servers: rawArgs,
	}
	ctx := cmd.GetContext(rawArgs)
	return commands.RunRestart(ctx, args)
}
