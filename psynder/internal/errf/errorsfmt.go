package errf

import "fmt"

func WithKindAndCause(desc string, kind int, cause error) string {
	return fmt.Sprintf("%v error with kind=%v caused by: %v", desc, kind, cause)
}

func WithCause(desc string, cause error) string {
	return fmt.Sprintf("%v error caused by: %v", desc, cause)
}
