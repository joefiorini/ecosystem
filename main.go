package main

import (
  "os"
  "fmt"
  "log"
  "io/ioutil"
  "github.com/mitchellh/cli"
)

func main() {
  os.Exit(run())
}

func run() int {
  log.SetOutput(ioutil.Discard)

  args := os.Args[1:]
  for _, arg := range args {
    if arg == "-v" || arg == "--version" {
      newArgs := make([]string, len(args)+1)
      newArgs[0] = "version"
      copy(newArgs[1:], args)
      args = newArgs
      break
    }
  }

  cli := &cli.CLI{
    Args: args,
    Commands: Commands,
  }

  exitCode, err := cli.Run()

  if err != nil {
    fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
    return 1
  }

  return exitCode
}
