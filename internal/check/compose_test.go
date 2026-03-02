package check

import (
	"context"
	"errors"
	"testing"
)

func TestComposeCheck_AllRunning(t *testing.T) {
	c := &ComposeCheck{runner: func() ([]byte, error) {
		return []byte(`{"Name":"app-db-1","State":"running"}` + "\n" +
			`{"Name":"app-web-1","State":"running"}` + "\n"), nil
	}}
	result := c.Run(context.Background())
	if result.Status != StatusPass {
		t.Errorf("expected pass, got %v: %s", result.Status, result.Message)
	}
}

func TestComposeCheck_SomeStopped(t *testing.T) {
	c := &ComposeCheck{runner: func() ([]byte, error) {
		return []byte(`{"Name":"app-db-1","State":"running"}` + "\n" +
			`{"Name":"app-web-1","State":"exited"}` + "\n"), nil
	}}
	result := c.Run(context.Background())
	if result.Status != StatusFail {
		t.Errorf("expected fail, got %v", result.Status)
	}
}

func TestComposeCheck_CommandFails(t *testing.T) {
	c := &ComposeCheck{runner: func() ([]byte, error) {
		return nil, errors.New("docker not found")
	}}
	result := c.Run(context.Background())
	if result.Status != StatusFail {
		t.Errorf("expected fail, got %v", result.Status)
	}
}

func TestComposeCheck_NoServices(t *testing.T) {
	c := &ComposeCheck{runner: func() ([]byte, error) {
		return []byte(""), nil
	}}
	result := c.Run(context.Background())
	if result.Status != StatusPass {
		t.Errorf("expected pass, got %v: %s", result.Status, result.Message)
	}
}
