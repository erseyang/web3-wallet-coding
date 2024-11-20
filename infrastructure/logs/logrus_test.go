package logs

import "testing"

func TestLogger(t *testing.T) {
	InitLogrus()
	logger.Info("This is a test")
}
