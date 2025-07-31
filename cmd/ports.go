package cmd

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
)

var portsCmd = &cobra.Command{
	Use:   "ports [container-name]",
	Short: "List port mappings for a container",
	Long:  "List port mappings for a container in a clean format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]

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

		fmt.Printf("Port mappings for container: %s\n", containerName)
		fmt.Println(strings.Repeat("-", 50))

		if len(containerInfo.NetworkSettings.Ports) == 0 {
			fmt.Println("No port mappings found")
			return
		}

		// Sort ports for consistent output
		var ports []string
		for port := range containerInfo.NetworkSettings.Ports {
			ports = append(ports, string(port))
		}
		sort.Strings(ports)

		fmt.Printf("%-15s %-15s %s\n", "CONTAINER PORT", "HOST PORT", "PROTOCOL")
		fmt.Println(strings.Repeat("-", 50))

		for _, portStr := range ports {
			bindings := containerInfo.NetworkSettings.Ports[nat.Port(portStr)]
			if len(bindings) > 0 {
				for _, binding := range bindings {
					hostPort := binding.HostPort
					if hostPort == "" {
						hostPort = "N/A"
					}
					protocol := strings.Split(portStr, "/")[1]
					containerPort := strings.Split(portStr, "/")[0]
					fmt.Printf("%-15s %-15s %s\n", containerPort, hostPort, protocol)
				}
			} else {
				protocol := strings.Split(portStr, "/")[1]
				containerPort := strings.Split(portStr, "/")[0]
				fmt.Printf("%-15s %-15s %s\n", containerPort, "Not published", protocol)
			}
		}

		// Show network information
		fmt.Println("\nNetwork Information:")
		fmt.Println(strings.Repeat("-", 30))
		for networkName, network := range containerInfo.NetworkSettings.Networks {
			fmt.Printf("Network: %s\n", networkName)
			fmt.Printf("  IP Address: %s\n", network.IPAddress)
			fmt.Printf("  Gateway: %s\n", network.Gateway)
			fmt.Printf("  Mac Address: %s\n", network.MacAddress)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(portsCmd)
}
