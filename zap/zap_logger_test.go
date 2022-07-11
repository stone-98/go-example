package zap

import (
	"go.uber.org/zap"
	"net/http"
	"testing"
)

var logger *zap.Logger

func TestZapLogger(t *testing.T) {
	initLogger()
	defer logger.Sync()
	simpleHttpGetByLogger("http://www.baidu.com")
	simpleHttpGetByLogger("192.168.0.188/hello")
}

func initLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGetByLogger(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("error fetching url...", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("Success...", zap.String("statusCode", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}
}
