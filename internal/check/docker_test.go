package check

import (
	"context"
	"errors"
	"testing"
)

func TestDockerDaemonCheck_Pass(t *testing.T) {
	c := &DockerDaemonCheck{runner: func() error { return nil }}
	result := c.Run(context.Background())
	if result.Status != StatusPass {
		t.Errorf("expected pass, got %v: %s", result.Status, result.Message)
	}
}

func TestDockerDaemonCheck_Fail(t *testing.T) {
	c := &DockerDaemonCheck{runner: func() error { return errors.New("daemon not running") }}
	result := c.Run(context.Background())
	if result.Status != StatusFail {
		t.Errorf("expected fail, got %v", result.Status)
	}
}
