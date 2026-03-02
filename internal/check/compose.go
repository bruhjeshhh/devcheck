package check

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type ComposeCheck struct {
	runner func() ([]byte, error)
}

func (c *ComposeCheck) Name() string {
	return "Docker Compose services running"
}

type composeService struct {
	Name  string `json:"Name"`
	State string `json:"State"`
}

func (c *ComposeCheck) Run(_ context.Context) Result {
	run := c.runner
	if run == nil {
		run = func() ([]byte, error) {
			return exec.Command("docker", "compose", "ps", "--format", "json").Output()
		}
	}

	out, err := run()
	if err != nil {
		return Result{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "could not run docker compose ps",
			Fix:     "make sure Docker is running and you are in the project directory",
		}
	}

	// docker compose ps --format json outputs one JSON object per line
	var stopped []string
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var svc composeService
		if err := json.Unmarshal([]byte(line), &svc); err != nil {
			continue
		}
		if svc.State != "running" {
			stopped = append(stopped, svc.Name)
		}
	}

	if len(stopped) > 0 {
		return Result{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: fmt.Sprintf("services not running: %s", strings.Join(stopped, ", ")),
			Fix:     "run docker compose up -d to start them",
		}
	}

	return Result{
		Name:    c.Name(),
		Status:  StatusPass,
		Message: "all services are running",
	}
}
