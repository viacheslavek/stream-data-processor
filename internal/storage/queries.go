package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func FileStats() error {
	// TODO: получать характеристики файла
	return nil
}

func GetDockerMemoryUsage(containerName string) (uint64, error) {
	stat, err := DockerStats(containerName)
	if err != nil {
		return 0, fmt.Errorf("failed get docker memory %w", err)
	}
	return stat.MemoryStats.Usage, nil
}

func DockerStats(containerName string) (types.Stats, error) {
	cli, errNC := client.NewClientWithOpts(client.FromEnv)
	if errNC != nil {
		return types.Stats{}, fmt.Errorf("failed to create docker client %w", errNC)
	}

	containerInfo, errGC := getContainerInfo(cli, containerName)
	if errGC != nil {
		return types.Stats{}, fmt.Errorf("failed to get container info %w", errGC)
	}

	containerID := containerInfo.ID

	stats, errCS := cli.ContainerStats(context.Background(), containerID, false)
	if errCS != nil {
		return types.Stats{}, fmt.Errorf("failed to get container stats %w", errCS)
	}
	defer func() {
		_ = stats.Body.Close()
	}()

	var stat types.Stats
	if err := json.NewDecoder(stats.Body).Decode(&stat); err != nil {
		return types.Stats{}, fmt.Errorf("failed to decod stats %w", err)
	}

	return stat, nil
}

func getContainerInfo(cli *client.Client, containerName string) (types.ContainerJSON, error) {
	filter := filters.NewArgs()
	filter.Add("name", containerName)

	containers, errCL := cli.ContainerList(context.Background(), container.ListOptions{Filters: filter})
	if errCL != nil {
		return types.ContainerJSON{}, fmt.Errorf("failed to get containers list %w", errCL)
	}

	if len(containers) == 0 {
		return types.ContainerJSON{}, fmt.Errorf("container with name %s not found", containerName)
	}

	containerID := containers[0].ID

	containerInfo, errCI := cli.ContainerInspect(context.Background(), containerID)
	if errCI != nil {
		return types.ContainerJSON{}, fmt.Errorf("failed inspect container %w", errCI)
	}

	return containerInfo, nil
}
