package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kubeless/kubeless/pkg/utils"
	"github.com/spf13/cobra"
)

var autoscaleCreateCmd = &cobra.Command{
	Use:   "create <name> FLAG",
	Short: "automatically scale function based on monitored metrics",
	Long:  `automatically scale function based on monitored metrics`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Fatal("Need exactly one argument - function name")
		}
		funcName := args[0]

		min, err := cmd.Flags().GetInt32("min")
		if err != nil {
			logrus.Fatal(err)
		} else if min == 0 {
			logrus.Fatalf("--min can't be 0")
		}
		max, err := cmd.Flags().GetInt32("max")
		if err != nil {
			logrus.Fatal(err)
		} else if min == 0 {
			logrus.Fatalf("--max can't be 0")
		}
		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			logrus.Fatal(err.Error())
		}

		client := utils.GetClientOutOfCluster()

		err = utils.CreateAutoscale(client, funcName, ns, min, max)
		if err != nil {
			logrus.Fatalf("Can't create autoscale: %v", err)
		}
	},
}

func init() {
	autoscaleCreateCmd.Flags().Int32("min", 0, "minimum number of replicas")
	autoscaleCreateCmd.Flags().Int32("max", 0, "maximum number of replicas")
}
