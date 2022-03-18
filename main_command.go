package main

//import (
//	"fmt"
//	"github.com/sirupsen/logrus"
//	"github.com/urfave/cli"
//)
//
//var runCommand = &cli.Command{
//	Name:  "run",
//	Usage: `Create a container with namespace and cgroups limit i-docker run -ti [command]`,
//	Flags: []cli.Flag{
//		&cli.BoolFlag{
//			Name:  "ti",
//			Usage: "enable tty",
//		},
//	},
//	Action: func(context *cli.Context) error {
//		if context.Args().Len() < 1 {
//			return fmt.Errorf("Missing container command")
//		}
//		cmd := context.Args().Get(0)
//		tty := context.Bool("ti")
//		fmt.Sprintf("cmd is %v, tty is %v\n", cmd, tty)
//		return nil
//	},
//}
//
//var initCommand = &cli.Command{
//	Name:  "init",
//	Usage: "Init container process run user's process in container. Do not call it outside",
//	Action: func(context *cli.Context) error {
//		logrus.Infof("init come on")
//		cmd := context.Args().Get(0)
//		logrus.Infof("command %s", cmd)
//
//		return nil
//	},
//}
