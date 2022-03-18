package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"i-docker/container"
	"i-docker/run"
	"os"
)

const usage = `mydocker is a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves
			   Enjoy it, just for fun.`

var (
	tty     bool
	rootCmd = &cobra.Command{
		Use:  "i-docker",
		Long: usage,
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf("failed to run, error is %v", err)
		os.Exit(1)
	}
	return
}

func init() {
	rootCmd.AddCommand(newInitCommand())
	rootCmd.AddCommand(newRunCommand())
}

func newInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "init",
		Long: "Init container process run user's process in container. Do not call it outside",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("init come on")
			command := args[0]
			logrus.Infof("command %s", command)
			err := container.RunContainerInitProcess(command, nil)
			if err != nil {
				logrus.Errorf("failed to run container init process, error is %v", err)
			}
		},
	}
}

func newRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "run",
		Long: `Create a container with namespace and cgroups limit i-docker run -ti [command]`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Errorf("Missing container command")
			}
			command := args[0]
			logrus.Infof("cmd is %v, tty is %v\n", command, tty)
			run.Run(tty, command)
		},
	}
	cmd.Flags().BoolVar(&tty, "tty", false, "enable tty")
	return cmd
}
