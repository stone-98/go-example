package zap

import (
	"go.uber.org/zap"
	"net/http"
	"testing"
)

var sugarLogger *zap.SugaredLogger

func TestZapSugarLogger(t *testing.T) {
	initSugarLogger()
	defer sugarLogger.Sync()
	simpleHttpGetBySugarLogger("http://www.baidu.com")
	simpleHttpGetBySugarLogger("192.168.113")
}

func initSugarLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGetBySugarLogger(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
