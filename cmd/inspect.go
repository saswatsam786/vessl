package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [container-name]",
	Short: "Display detailed information on a container",
	Long:  "Display detailed information on a container in JSON format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		jsonOutput, _ := cmd.Flags().GetBool("json")

		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		containerInfo, err := client.ContainerInspect(context.Background(), containerName)
		if err != nil {
			fmt.Printf("Error inspecting container: %v\n", err)
			return
		}

		if jsonOutput {
			jsonData, err := json.MarshalIndent(containerInfo, "", "  ")
			if err != nil {
				fmt.Printf("Error marshaling JSON: %v\n", err)
				return
			}
			fmt.Println(string(jsonData))
		} else {
			fmt.Printf("Container ID: %s\n", containerInfo.ID)
			fmt.Printf("Name: %s\n", containerInfo.Name)
			fmt.Printf("Image: %s\n", containerInfo.Config.Image)
			fmt.Printf("Status: %s\n", containerInfo.State.Status)
			fmt.Printf("Created: %s\n", containerInfo.Created)
		}
	},
}

func init() {
	inspectCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
	rootCmd.AddCommand(inspectCmd)
}
