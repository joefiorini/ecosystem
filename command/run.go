package command

import (
  "fmt"
  "flag"
  "io/ioutil"
  "strings"
  "github.com/mitchellh/cli"
  "github.com/joefiorini/ecosystem/docker"
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

// docker run -v /mnt/hgfs/{CURRENT_DIR}:/usr/local/src/app -link {c.LINK_1} -link {c.LINK_2} -name {c.NAME} -u {c.USER} -w /usr/local/src/app {c.TEMPLATE} {c.PREFIX|/usr/bin/zsh -ic} {CMD}
func (c *RunCommand) Run(args []string) int {
  cmdFlags := flag.NewFlagSet("run", flag.ContinueOnError)
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

  argFileContents,err := ioutil.ReadFile(".ecosystem")

  if err != nil {
    argFileContents = []byte("")
  }

  cmdString := strings.TrimSpace(strings.Join(cmd, " "))
  runString := strings.TrimSpace(string(argFileContents))
  runArgs := strings.Split(runString, " ")
  c.Ui.Output(fmt.Sprintf("run %s %s", runString, cmdString))

  client := new(docker.Docker)
  client.Addr = "triforce.local"
  client.Port = "4243"
  client.Debug = c.Debug

  client.Run(cmdString, runArgs...)

  return 0
}

func (c *RunCommand) Synopsis() string {
        return "Runs a command within an ecosystem"
}
