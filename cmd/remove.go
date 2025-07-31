package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [container ID]",
	Short: "Remove a docker container",
	Long:  "Remove a container with a specified ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		force, _ := cmd.Flags().GetBool("force")
		removeOptions := container.RemoveOptions{
			Force: force,
		}

		err = client.ContainerRemove(context.Background(), containerName, removeOptions)
		if err != nil {
			fmt.Printf("Error removing container: %v\n", err)
			return
		}
		fmt.Printf("Container %s removed successfully\n", containerName)
	},
}

func init() {
	// add a force flag to the remove command
	removeCmd.Flags().BoolP("force", "f", false, "Force removal of the container")
	rootCmd.AddCommand(removeCmd)
}
