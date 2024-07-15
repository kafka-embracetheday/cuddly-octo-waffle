package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	initLogger()
}

func initLogger() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	Logger.SetReportCaller(true)
}

func Debug(args ...interface{}) {
	writeLog(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	writeLog(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	writeLog(logrus.WarnLevel, args...)
}

func Error(args ...interface{}) {
	writeLog(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	writeLog(logrus.FatalLevel, args...)
}

func Panic(args ...interface{}) {
	writeLog(logrus.PanicLevel, args...)
}

func writeLog(level logrus.Level, args ...interface{}) {

	logFileName := time.Now().Format("2006-01-02") + ".log"
	logDirPath := "logs"
	logPath := filepath.Join(logDirPath, logFileName)

	// 如果 logs 目录不存在,则创建
	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		err := os.MkdirAll(logDirPath, 0755)
		if err != nil {
			Logger.Errorf("Failed to create log directory: %v", err)
			return
		}
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// 如果打开日志文件失败,则回退到控制台输出
		Logger.SetOutput(os.Stdout)
		Logger.Errorf("Failed to open log file: %v", err)
	} else {
		defer file.Close()
		Logger.SetOutput(file)

		switch level {
		case logrus.DebugLevel:
			Logger.Debugln(args...)
		case logrus.InfoLevel:
			Logger.Infoln(args...)
		case logrus.WarnLevel:
			Logger.Warnln(args...)
		case logrus.ErrorLevel:
			Logger.Errorln(args...)
		case logrus.FatalLevel:
			Logger.Fatalln(args...)
		case logrus.PanicLevel:
			Logger.Panicln(args...)
		}
	}
}
