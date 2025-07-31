package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List images",
	Long:  "List all Docker images",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		images, err := client.ImageList(context.Background(), types.ImageListOptions{})
		if err != nil {
			fmt.Printf("Error listing images: %v\n", err)
			return
		}

		fmt.Printf("%-12s %-20s %-15s %s\n", "IMAGE ID", "REPOSITORY", "TAG", "SIZE")
		fmt.Println(strings.Repeat("-", 70))

		for _, image := range images {
			repo := "none"
			tag := "none"
			if len(image.RepoTags) > 0 {
				parts := strings.Split(image.RepoTags[0], ":")
				if len(parts) == 2 {
					repo = parts[0]
					tag = parts[1]
				}
			}
			fmt.Printf("%-12s %-20s %-15s %s\n",
				image.ID[:12], repo, tag, formatSize(image.Size))
		}
	},
}

func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}
