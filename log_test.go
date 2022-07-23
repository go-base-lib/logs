package logs

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLogRecord(t *testing.T) {
	record, err := NewLogRecord(&Config{
		PathConfig: &PathConfig{
			LogPath:      filepath.Join("logs", "log.txt.%Y%m%d%H%M"),
			LinkName:     filepath.Join("logs", "log.txt"),
			RotationTime: 5 * time.Minute,
			RotationSize: 100,
		},
		Formatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		CurrentLevel: logrus.DebugLevel,
	})
	if err != nil {
		t.Error("创建日志对象出错 => ", record)
		return
	}

	record.Info("我是info级别的日志")

	_ = os.RemoveAll("logs")

}
