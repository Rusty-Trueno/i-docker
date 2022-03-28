package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"i-docker/cgroups/subsystems"
	"i-docker/container"
	"i-docker/run"
	"os"
)

const usage = `mydocker is a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves
			   Enjoy it, just for fun.`

var (
	tty      bool
	m        string
	cpushare string
	cpuset   string
	rootCmd  = &cobra.Command{
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
	pflag.BoolVar(&tty, "ti", false, "enable tty")
	pflag.StringVar(&m, "m", "", "memory limit")
	pflag.StringVar(&cpushare, "cpushare", "", "cpushare limit")
	pflag.StringVar(&cpuset, "cpuset", "", "cpuset limit")
	rootCmd.AddCommand(newInitCommand())
	rootCmd.AddCommand(newRunCommand())
	rootCmd.DisableSuggestions = true
}

func newInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "init",
		Long: "Init container process run user's process in container. Do not call it outside",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("init come on")
			err := container.RunContainerInitProcess()
			if err != nil {
				logrus.Errorf("failed to run container init process, error is %v", err)
			}
		},
	}
	return cmd
}

func newRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "run",
		Long: `Create a container with namespace and cgroups limit i-docker run -ti [command]`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Errorf("Missing container command")
				return
			}
			var cmdArray []string
			for _, arg := range args {
				cmdArray = append(cmdArray, arg)
			}
			tty := tty
			resConf := &subsystems.ResourceConfig{
				MemoryLimit: m,
				CpuSet:      cpuset,
				CpuShare:    cpushare,
			}

			run.Run(tty, cmdArray, resConf)
		},
	}
	return cmd
}
