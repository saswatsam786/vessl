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

var pullCmd = &cobra.Command{
	Use:   "pull [image-name]",
	Short: "Pull an image from a registry",
	Long:  "Pull an image from a registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]

		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		fmt.Printf("Pulling image: %s\n", imageName)
		reader, err := client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
		if err != nil {
			fmt.Printf("Error pulling image: %v\n", err)
			return
		}
		defer reader.Close()

		io.Copy(os.Stdout, reader)
		fmt.Printf("Successfully pulled image: %s\n", imageName)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
