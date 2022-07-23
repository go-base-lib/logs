package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	PathConfig   *PathConfig
	Formatter    logrus.Formatter
	Hooks        []logrus.Hook
	CurrentLevel logrus.Level
	// 设置是否添加文件名和方法信息
	ReportCaller bool
}

type PathConfig struct {
	LogPath       string
	LinkName      string
	MaxAge        time.Duration
	RotationTime  time.Duration
	RotationCount uint
	RotationSize  int64
	Loc           *time.Location
	Clock         rotatelogs.Clock
}

func (this *PathConfig) SettingPath(p string) *PathConfig {
	return this.SettingPathWithSuffix(p, ".%Y%m%d%H%M")
}

func (this *PathConfig) SettingPathWithSuffix(p, suffix string) *PathConfig {
	this.LogPath = p + suffix
	this.LinkName = p
	return this
}

func (this *PathConfig) getRotateConfig() (*rotatelogs.RotateLogs, error) {

	options := make([]rotatelogs.Option, 0)
	if this.RotationSize != 0 {
		options = append(options, rotatelogs.WithRotationSize(this.RotationSize))
	}

	if this.RotationCount != 0 {
		options = append(options, rotatelogs.WithRotationCount(this.RotationCount))
	}

	if this.RotationTime > 0 {
		options = append(options, rotatelogs.WithRotationTime(this.RotationTime))
	}

	if this.LinkName != "" {
		options = append(options, rotatelogs.WithLinkName(this.LinkName))
	}

	if this.MaxAge > 0 {
		options = append(options, rotatelogs.WithMaxAge(this.MaxAge))
	}

	if this.Loc != nil {
		options = append(options, rotatelogs.WithLocation(this.Loc))
	}

	if this.Clock != nil {
		options = append(options, rotatelogs.WithClock(this.Clock))
	}

	return rotatelogs.New(
		this.LogPath,
		options...,
	)

}
