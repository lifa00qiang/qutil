package QUtil

import (
	"go.uber.org/zap"
	"testing"
)

func TestGetLogger(t *testing.T) {

	logger := GetLogger()
	logger.Info("111", zap.String("123123", "123123"))
}
