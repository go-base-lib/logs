package logs

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Record struct {
	*logrus.Logger
	Config *Config
}

func (this *Record) Init() error {
	if this.Logger == nil {
		this.Logger = logrus.New()
	}

	if this.Config == nil {
		return nil
	}

	if this.Config.ReportCaller {
		this.Logger.SetReportCaller(this.Config.ReportCaller)
	}

	if this.Config.Formatter != nil {
		this.Logger.SetFormatter(this.Config.Formatter)
	}

	if this.Config.PathConfig != nil {
		pathConfig := this.Config.PathConfig
		config, err := pathConfig.getRotateConfig()
		if err != nil {
			return err
		}
		hook := lfshook.NewHook(lfshook.WriterMap{
			logrus.TraceLevel: config,
			logrus.DebugLevel: config,
			logrus.InfoLevel:  config,
			logrus.WarnLevel:  config,
			logrus.ErrorLevel: config,
			logrus.FatalLevel: config,
			logrus.PanicLevel: config,
		}, this.Logger.Formatter)
		this.Logger.AddHook(hook)
	}

	if len(this.Config.Hooks) > 0 {
		for _, v := range this.Config.Hooks {
			this.Logger.AddHook(v)
		}
	}

	this.Logger.SetLevel(this.Config.CurrentLevel)

	return nil
}

var defaultRecord *Record

func NewLogRecord(config *Config) (*Record, error) {
	record := &Record{
		Config: config,
	}

	return record, record.Init()
}

func InitDefault(config *Config) error {
	record, err := NewLogRecord(config)
	if err != nil {
		return err
	}
	defaultRecord = record
	return nil
}

func init() {
	config := &Config{
		PathConfig: &PathConfig{},
	}
	logger := logrus.New()
	defaultRecord = &Record{
		Logger: logger,
		Config: config,
	}
}
