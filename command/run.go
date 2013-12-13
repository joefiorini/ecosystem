package command

import (
  "fmt"
  "flag"
  "strings"
  "github.com/mitchellh/cli"
)

type RunCommand struct {
  Ui cli.Ui
}

func (c *RunCommand) Help() string {
  helpText := `
Usage: eco run COMMAND ...

  Runs a command with arguments inside an ecosystem.
`
  return strings.TrimSpace(helpText)
}

func (c *RunCommand) Run(args []string) int {
  cmdFlags := flag.NewFlagSet("exec", flag.ContinueOnError)
  cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
  if err := cmdFlags.Parse(args); err != nil {
    return 1
  }

  cmd := cmdFlags.Args()
  if len(cmd) == 0 {
    c.Ui.Error("You must specify a command.")
    c.Ui.Error("")
    c.Ui.Error(c.Help())
    return 1
  }

  c.Ui.Output(fmt.Sprintf("exec %a", cmd))

  return 0
}

func (c *RunCommand) Synopsis() string {
        return "Runs a command within an ecosystem"
}
