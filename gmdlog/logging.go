package gmdlog

import (
	log "github.com/sirupsen/logrus"
	"os"
	"io"
)

type GmdLog struct {}


func (l GmdLog) InitializeLogger(isProduction bool) {
	log.SetFormatter(&log.JSONFormatter{})

	logFile, err := os.Open("noteclerk.log")
	if err != nil {
		panic("Failed to access log file.")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	if isProduction {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
}

func (l GmdLog) WithField(key string, value interface{}) *log.Entry {
	return log.WithField(key, value)
}

func (l GmdLog) WithFields(fields log.Fields) *log.Entry {
	return log.WithFields(fields)
}

func (l GmdLog) WithError(err error) *log.Entry {
	return log.WithError(err)
}

func (l GmdLog) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}

func (l GmdLog) Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func (l GmdLog) Printf(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (l GmdLog) Warnf(format string, args ...interface{}) {
	log.Warnf(format, args)
}

func (l GmdLog) Warningf(format string, args ...interface{}) {
	log.Warningf(format, args)
}

func (l GmdLog) Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}

func (l GmdLog) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

func (l GmdLog) Panicf(format string, args ...interface{}) {
	log.Panicf(format, args)
}

func (l GmdLog) Debug(args ...interface{}) {
	log.Debug(args)
}

func (l GmdLog) Info(args ...interface{}) {
	log.Info(args)
}

func (l GmdLog) Print(args ...interface{}) {
	log.Print(args)
}

func (l GmdLog) Warn(args ...interface{}) {
	log.Warn(args)
}

func (l GmdLog) Warning(args ...interface{}) {
	log.Warning(args)
}

func (l GmdLog) Error(args ...interface{}) {
	log.Error(args)
}

func (l GmdLog) Fatal(args ...interface{}) {
	log.Fatal(args)
}

func (l GmdLog) Panic(args ...interface{}) {
	log.Panic(args)
}

func (l GmdLog) Debugln(args ...interface{}) {
	log.Debugln(args)
}

func (l GmdLog) Infoln(args ...interface{}) {
	log.Infoln(args)
}

func (l GmdLog) Println(args ...interface{}) {
	log.Println(args)
}

func (l GmdLog) Warnln(args ...interface{}) {
	log.Warnln(args)
}

func (l GmdLog) Warningln(args ...interface{}) {
	log.Warningln(args)
}

func (l GmdLog) Errorln(args ...interface{}) {
	log.Errorln(args)
}

func (l GmdLog) Fatalln(args ...interface{}) {
	log.Fatalln(args)
}

func (l GmdLog) Panicln(args ...interface{}) {
	log.Panicln(args)
}
