package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"io"
)

type GmdLog struct {}


func (l GmdLog) InitializeLogger(logPath string) {
	flags := os.O_APPEND | os.O_WRONLY
	logFile, err := os.OpenFile(logPath, flags, os.ModeAppend)
	if err != nil {
		log.Fatalf("Cannot access or locate log file at location %v.", logPath)
	}

	logrus.SetLevel(logrus.InfoLevel)
	var writer io.Writer = os.Stdout
	if Environment != "production" {
		writer = io.MultiWriter(os.Stdout, logFile)
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetFormatter(&logrus.TextFormatter{})

	logrus.SetOutput(writer)
}

func (l GmdLog) WithField(key string, value interface{}) *logrus.Entry {
	return logrus.WithField(key, value)
}

func (l GmdLog) WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.WithFields(fields)
}

func (l GmdLog) WithError(err error) *logrus.Entry {
	return logrus.WithError(err)
}

func (l GmdLog) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args)
}

func (l GmdLog) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args)
}

func (l GmdLog) Printf(format string, args ...interface{}) {
	logrus.Printf(format, args)
}

func (l GmdLog) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args)
}

func (l GmdLog) Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args)
}

func (l GmdLog) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args)
}

func (l GmdLog) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args)
}

func (l GmdLog) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args)
}

func (l GmdLog) Debug(args ...interface{}) {
	logrus.Debug(args)
}

func (l GmdLog) Info(args ...interface{}) {
	logrus.Info(args)
}

func (l GmdLog) Print(args ...interface{}) {
	logrus.Print(args)
}

func (l GmdLog) Warn(args ...interface{}) {
	logrus.Warn(args)
}

func (l GmdLog) Warning(args ...interface{}) {
	logrus.Warning(args)
}

func (l GmdLog) Error(args ...interface{}) {
	logrus.Error(args)
}

func (l GmdLog) Fatal(args ...interface{}) {
	logrus.Fatal(args)
}

func (l GmdLog) Panic(args ...interface{}) {
	logrus.Panic(args)
}

func (l GmdLog) Debugln(args ...interface{}) {
	logrus.Debugln(args)
}

func (l GmdLog) Infoln(args ...interface{}) {
	logrus.Infoln(args)
}

func (l GmdLog) Println(args ...interface{}) {
	logrus.Println(args)
}

func (l GmdLog) Warnln(args ...interface{}) {
	logrus.Warnln(args)
}

func (l GmdLog) Warningln(args ...interface{}) {
	logrus.Warningln(args)
}

func (l GmdLog) Errorln(args ...interface{}) {
	logrus.Errorln(args)
}

func (l GmdLog) Fatalln(args ...interface{}) {
	logrus.Fatalln(args)
}

func (l GmdLog) Panicln(args ...interface{}) {
	logrus.Panicln(args)
}
