package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [container-name]",
	Short: "Fetch the logs of a container",
	Long:  "Fetch the logs of a container with optional follow and tail flags",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		follow, _ := cmd.Flags().GetBool("follow")
		tail, _ := cmd.Flags().GetString("tail")

		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		options := types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     follow,
		}

		if tail != "" {
			options.Tail = tail
		}

		logs, err := client.ContainerLogs(context.Background(), containerName, options)
		if err != nil {
			fmt.Printf("Error fetching logs: %v\n", err)
			return
		}
		defer logs.Close()

		io.Copy(os.Stdout, logs)
	},
}

func init() {
	logsCmd.Flags().BoolP("follow", "f", false, "Follow log output")
	logsCmd.Flags().StringP("tail", "t", "", "Number of lines to show from the end of the logs")
	rootCmd.AddCommand(logsCmd)
}
