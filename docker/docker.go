package docker

import (
  "os"
  "os/exec"
  "fmt"
  "strings"
)

type Docker struct {
  Addr string
  Port string
  Debug bool
}

func (docker *Docker) buildCommand(command string, arg ...string) (*exec.Cmd) {
  args := make([]string, 1)

  if docker.Addr != "" {
    newArgs := make([]string, len(args)+1)
    copy(newArgs, args)

    host := fmt.Sprintf("-H=tcp://%s", docker.Addr)

    if docker.Port != "" {
      host = fmt.Sprintf("%s:%s", host, docker.Port)
    }

    newArgs[0] = host
    args = newArgs
  }

  if args[0] != "" {
    args[1] = command
  } else {
    args[0] = command
  }

  args = append(args, arg...)
  cmd := exec.Command("/usr/local/bin/docker", args...)
  return cmd
}

func (docker *Docker) Exec(command string, arg ...string) (error) {
  cmd := docker.buildCommand(command, arg...)

  if docker.Debug {
    fmt.Println(cmd.Args)
  }

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Start()


  if err != nil {
    return err
  }

  err = cmd.Wait()

  return err
}

func (docker *Docker) Run(cmd []string, arg ...string) error {
  args := append(arg, strings.Join(cmd, " "))
  return docker.Exec("run", args...)
}

