package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all docker containers",
	Long:  "List all docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Failed to create docker client: %v\n", err)
			return
		}
		defer client.Close()

		containers, err := client.ContainerList(context.Background(), container.ListOptions{All: true})
		if err != nil {
			fmt.Printf("Failed to list containers: %v\n", err)
			return
		}

		fmt.Printf("%-12s %-20s %-30s %-15s %s\n", "CONTAINER ID", "IMAGE", "COMMAND", "CREATED", "STATUS")
		fmt.Println(strings.Repeat("-", 100))
		for _, container := range containers {
			createdStr := time.Unix(container.Created, 0).Format("2006-01-02 15:04:05")
			fmt.Printf("%-12s %-20s %-30s %-15s %s\n", container.ID[:12], container.Image, container.Names[0], createdStr, container.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
