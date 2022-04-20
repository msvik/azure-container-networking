package platform

import (
	"testing"
	"time"
)

// Default timeout in ExecuteCommand is 10 seconds so test should fail with 12 seconds execution time
func TestExecuteCommandTimeout(t *testing.T) {
	const timeout = 10 * time.Second
	client := NewExecClientTimeout(timeout)

	_, err := client.ExecuteCommand("sleep 12")
	if err == nil {
		t.Errorf("TestExecuteCommandTimeout should have returned timeout error")
	}
	t.Logf("%s", err.Error())
}

// Default timeout in ExecuteCommand is 10 seconds so test should pass with 8 seconds execution time
func TestExecuteCommandNoTimeout(t *testing.T) {
	const timeout = 10 * time.Second
	client := NewExecClientTimeout(timeout)

	_, err := client.ExecuteCommand("sleep 8")
	if err != nil {
		t.Errorf("TestExecuteCommandNoTimeout failed with error %v", err)
	}
}
