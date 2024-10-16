package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type LogConf struct {
	Level      int
	Format     string
	Output     string
	Caller     bool
	OutputFile string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
type Hook = logrus.Hook

func MustSetup(c *LogConf) {
	SetFormatter(c.Format)
	SetLevel(c.Level)
	SetReportCaller(c.Caller)
	switch c.Output {
	case "stdout":
		SetOutput(os.Stdout)
	case "stderr":
		SetOutput(os.Stderr)
	case "file":
		lumber := &lumberjack.Logger{
			Filename:   c.OutputFile,
			MaxSize:    c.MaxSize,
			MaxBackups: c.MaxBackups,
			MaxAge:     c.MaxAge,
			Compress:   c.Compress,
		}
		SetOutput(lumber)
	default:
		SetOutput(os.Stdout)
	}
}
func SetFormatter(formatter string) {
	switch formatter {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	}
}

func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

func SetLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}
func AddHook(hook Hook) {
	logrus.AddHook(hook)
}

// 显示行号
func SetReportCaller(reportCaller bool) {
	logrus.SetReportCaller(reportCaller)
}
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return logrus.WithFields(fields)
}

// Define logrus alias
var (
	Tracef  = logrus.Tracef
	Debugf  = logrus.Debugf
	Infof   = logrus.Infof
	Info    = logrus.Info
	Warnf   = logrus.Warnf
	Error   = logrus.Error
	Errorf  = logrus.Errorf
	Fatalf  = logrus.Fatalf
	Fatalln = logrus.Fatalln
	Panicf  = logrus.Panicf
	Printf  = logrus.Printf
	Println = logrus.Println
)
