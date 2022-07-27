package clierrors

import "fmt"

type CommandFailedError struct {
	FailedCount int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (err CommandFailedError) ErrorCode() int {
	return min(err.FailedCount, 99)
}

func (err CommandFailedError) Error() string {
	return fmt.Sprintf("Command execution failed for %d resource(s)", err.FailedCount)
}
