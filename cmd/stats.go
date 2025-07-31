package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats [container-name]",
	Short: "Display a live stream of container resource usage statistics",
	Long:  "Display a live stream of container resource usage statistics (CPU, memory, network)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			fmt.Printf("Error creating docker client: %v\n", err)
			return
		}
		defer client.Close()

		if len(args) == 0 {
			// Show stats for all running containers
			containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
			if err != nil {
				fmt.Printf("Error listing containers: %v\n", err)
				return
			}

			fmt.Printf("%-12s %-20s %-8s %-8s %-8s %-8s %-8s %s\n",
				"CONTAINER", "NAME", "CPU %", "MEM USAGE", "MEM %", "NET I/O", "BLOCK I/O", "PIDS")
			fmt.Println(strings.Repeat("-", 100))

			for _, container := range containers {
				if container.State == "running" {
					go displayContainerStats(client, container.ID, container.Names[0])
				}
			}

			// Keep the main thread alive
			select {}
		} else {
			// Show stats for specific container
			containerName := args[0]
			displayContainerStats(client, containerName, containerName)
		}
	},
}

func displayContainerStats(client *client.Client, containerID, containerName string) {
	stats, err := client.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		fmt.Printf("Error getting stats for %s: %v\n", containerName, err)
		return
	}
	defer stats.Body.Close()

	var stat types.StatsJSON
	decoder := json.NewDecoder(stats.Body)
	if err := decoder.Decode(&stat); err != nil {
		fmt.Printf("Error decoding stats: %v\n", err)
		return
	}

	// Calculate CPU percentage
	cpuPercent := calculateCPUPercent(&stat)

	// Calculate memory usage
	memUsage := stat.MemoryStats.Usage - stat.MemoryStats.Stats["cache"]
	memLimit := stat.MemoryStats.Limit
	memPercent := float64(memUsage) / float64(memLimit) * 100.0

	// Calculate network I/O
	netRx := stat.Networks["eth0"].RxBytes
	netTx := stat.Networks["eth0"].TxBytes

	// Calculate block I/O
	blockRead := stat.BlkioStats.IoServiceBytesRecursive[0].Value
	blockWrite := stat.BlkioStats.IoServiceBytesRecursive[1].Value

	fmt.Printf("%-12s %-20s %-8.2f %-8s %-8.2f %-8s %-8s %d\n",
		containerID[:12],
		containerName,
		cpuPercent,
		formatBytes(memUsage),
		memPercent,
		fmt.Sprintf("%s / %s", formatBytes(netRx), formatBytes(netTx)),
		fmt.Sprintf("%s / %s", formatBytes(blockRead), formatBytes(blockWrite)),
		stat.PidsStats.Current)
}

func calculateCPUPercent(stats *types.StatsJSON) float64 {
	cpuPercent := 0.0
	if stats.CPUStats.CPUUsage.TotalUsage != 0 {
		cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
		systemDelta := float64(stats.CPUStats.SystemUsage - stats.PreCPUStats.SystemUsage)
		if systemDelta > 0.0 && cpuDelta > 0.0 {
			cpuPercent = (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
		}
	}
	return cpuPercent
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%dB", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
