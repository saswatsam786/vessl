package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new docker container",
	Long:  "Create a new docker container with specified name and image",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter the container name:")
		containerName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}
		containerName = strings.TrimSpace(containerName)

		fmt.Print("Enter the image name:")
		imageName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}
		imageName = strings.TrimSpace(imageName)

		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		// Check locally if image exists
		imageList, err := client.ImageList(context.Background(), types.ImageListOptions{})
		if err != nil {
			fmt.Printf("Error listing images: %v\n", err)
			return
		}

		imageExists := false

		for _, image := range imageList {
			for _, imgTag := range image.RepoTags {
				if imgTag == imageName {
					imageExists = true
					break
				}
			}
			if imageExists {
				break
			}
		}

		if !imageExists {
			fmt.Printf("Image not found locally. Pulling from the Docker hub.....\n")
			out, err := client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
			if err != nil {
				fmt.Printf("Error pulling image: %v\n", err)
				return
			}
			defer out.Close()
			io.Copy(os.Stdout, out)
		}

		container, err := client.ContainerCreate(context.Background(), &container.Config{
			Image: imageName,
		}, nil, nil, nil, containerName)
		if err != nil {
			fmt.Printf("Error creating container: %v\n", err)
			return
		}

		fmt.Printf("Container %s created with ID: %s\n", containerName, container.ID)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
