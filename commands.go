package main

import (
  "os"
  "github.com/mitchellh/cli"
  "github.com/joefiorini/ecosystem/command"
)

// Commands is the mapping of all the available Eco commands.
var Commands map[string]cli.CommandFactory

func init() {
        ui := &cli.BasicUi{Writer: os.Stdout}

        Commands = map[string]cli.CommandFactory{

                "run": func() (cli.Command, error) {
                  return &command.RunCommand{
                    Ui: ui,
                  }, nil
                },

                "version": func() (cli.Command, error) {
                        return &command.VersionCommand{
                                Revision:          GitCommit,
                                Version:           Version,
                                VersionPrerelease: VersionPrerelease,
                                Ui:                ui,
                        }, nil
                },
        }
}

