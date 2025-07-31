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

var execCmd = &cobra.Command{
	Use:   "exec [container-name] [command]",
	Short: "Run a command in a running container",
	Long:  "Run a command in a running container",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		command := args[1:]

		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		execConfig := types.ExecConfig{
			Cmd:          command,
			AttachStdout: true,
			AttachStderr: true,
			AttachStdin:  true,
			Tty:          true,
		}

		execResp, err := client.ContainerExecCreate(context.Background(), containerName, execConfig)
		if err != nil {
			fmt.Printf("Error creating exec: %v\n", err)
			return
		}

		resp, err := client.ContainerExecAttach(context.Background(), execResp.ID, types.ExecStartCheck{})
		if err != nil {
			fmt.Printf("Error attaching to exec: %v\n", err)
			return
		}
		defer resp.Close()

		// Copy output to stdout
		go func() {
			io.Copy(os.Stdout, resp.Reader)
		}()

		// Copy input from stdin
		io.Copy(resp.Conn, os.Stdin)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
