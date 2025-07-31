package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [container ID]",
	Short: "Stop a docker container",
	Long:  "Stop a container with a specified ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		err = client.ContainerStop(context.Background(), containerName, container.StopOptions{})
		if err != nil {
			fmt.Printf("Error stopping container: %v\n", err)
			return
		}
		fmt.Printf("Container %s stopped successfully\n", containerName)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
