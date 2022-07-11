package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"testing"
)

var customLogger *zap.SugaredLogger

func TestCustomLogger(t *testing.T) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	customLogger = logger.Sugar()
	simpleHttpGetBySugarLoggerByCustomLogger("http://www.baidu.com")
	simpleHttpGetBySugarLoggerByCustomLogger("192.168.0.113")
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func simpleHttpGetBySugarLoggerByCustomLogger(url string) {
	customLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		customLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		customLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
