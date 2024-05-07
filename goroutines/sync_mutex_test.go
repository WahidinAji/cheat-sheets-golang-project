package goroutines_test

import (
	"testing"
	goroutine "cheat-sheets-golang-project/goroutines"
)

func TestRunJob(t *testing.T) {
	task := goroutine.RunJob()
	if len(task) != 10 {
		t.Errorf("Expected 10, got %d", len(task))
	}
}
