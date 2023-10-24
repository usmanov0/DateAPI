package logger

import (
  "fmt"
  "io"
  "os"
  "path"
  "runtime"

  "github.com/sirupsen/logrus"

  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)


type writerHook struct {
  Writer    []io.Writer
  LogLevels []logrus.Level
}


func (hook *writerHook) Fire(entry *logrus.Entry) error {
  line, err := entry.String()
  if err != nil {
    return err
  }
  for _, w := range hook.Writer {
    _, err = w.Write([]byte(line))
  }
  return err
}


func (hook *writerHook) Levels() []logrus.Level {
  return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
  *logrus.Entry
}

func GetLogger() Logger {
  return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
  return Logger{l.WithField(k, v)}
}

func Init() {
  l := logrus.New()
  l.SetReportCaller(true)
  l.Formatter = &logrus.TextFormatter{
    CallerPrettyfier: func(f *runtime.Frame) (string, string) {
      filename := path.Base(f.File)
      return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
    },
    DisableColors: false,
    FullTimestamp: true,
  }

  err := os.MkdirAll("logs", 0755)

  if err != nil || os.IsExist(err) {
    panic("can't create log dir. no configured logging to files")
  } else {
    allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
      panic(fmt.Sprintf("[Message]: %s", err))
    }

    l.SetOutput(io.Discard) 

    l.AddHook(&writerHook{
      Writer:    []io.Writer{allFile, os.Stdout},
      LogLevels: logrus.AllLevels,
    })
  }

  l.SetLevel(logrus.TraceLevel)

  e = logrus.NewEntry(l)
}

type Field = zapcore.Field

var (
  Int = zap.Int
  String = zap.String
  Error = zap.Error
  Bool = zap.Bool
  Any = zap.Any
)

type LoggerImpl struct {
  zap *zap.Logger
}

func (l *LoggerImpl) Debug(msg string, fields ...Field) {
  l.zap.Debug(msg, fields...)
}

func (l *LoggerImpl) Info(msg string, fields ...Field) {
  l.zap.Info(msg, fields...)
}

func (l *LoggerImpl) Warn(msg string, fields ...Field) {
  l.zap.Warn(msg, fields...)
}

func (l *LoggerImpl) Error(msg string, fields ...Field) {
  l.zap.Error(msg, fields...)
}

func (l *LoggerImpl) Fatal(msg string, fields ...Field) {
  l.zap.Fatal(msg, fields...)
}