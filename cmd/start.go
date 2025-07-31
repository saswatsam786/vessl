package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [container ID]",
	Short: "Start a docker container",
	Long:  "Start a container with a specified ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}

		defer client.Close()

		err = client.ContainerStart(context.Background(), containerName, types.ContainerStartOptions{})
		if err != nil {
			fmt.Printf("Error starting container: %v\n", err)
			return
		}

		fmt.Printf("Container %s started successfully\n", containerName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
