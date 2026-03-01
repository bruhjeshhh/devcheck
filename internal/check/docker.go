package check

import (
	"context"
	"os/exec"
)

type DockerDaemonCheck struct {
	runner func() error
}

func (c *DockerDaemonCheck) Name() string {
	return "Docker daemon running"
}

func (c *DockerDaemonCheck) Run(_ context.Context) Result {
	run := c.runner
	if run == nil {
		run = func() error {
			return exec.Command("docker", "info").Run()
		}
	}

	if err := run(); err != nil {
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
