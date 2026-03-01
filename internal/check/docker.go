package check

import (
	"context"
	"os/exec"
)

type DockerDaemonCheck struct{}

func (c *DockerDaemonCheck) Name() string {
	return "Docker daemon running"
}

func (c *DockerDaemonCheck) Run(_ context.Context) Result {
	err := exec.Command("docker", "info").Run()
	if err != nil {
		return Result{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "Docker daemon is not running",
			Fix:     "start Docker and try again",
		}
	}
	return Result{
		Name:    c.Name(),
		Status:  StatusPass,
		Message: "Docker daemon is running",
	}
}
