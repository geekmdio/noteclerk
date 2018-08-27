package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"fmt"
	"strings"
	"github.com/pkg/errors"
)

// Inject preferred logger into the log global variable singleton. NOTE: at the time of development, this log singleton
// is dependent upon the logrus.Logger struct which implements the logrus.FieldLogger and logrus.StdLogger interfaces.
// custom loggers should be careful to implement these interfaces.
var log = logrus.New()

// The NOTECLERK_DATA environmental variable is retrieved fromm the OS and is used to determine what the root data
// directory is for configuration and log files.
var NoteClerkData = os.Getenv(DataRoot)

// The NOTECLERK_ENVIRONMENT environmental variable is retrieved from the OS and used to determine which configuration
// settings are to be loaded at runtime. New settings require a server restart.
var NoteClerkEnv = os.Getenv(Environment)

// Use the NOTECLERK_DATA path and the NOTECLERK_ENVIRONMENT environment type to generate a path to the preferred
// configuration file. Of note, it is suggested that this be limited to a user scope if possible. Manual manipulation of
// permissions may be required if moving configuration and logging files to system folders.
var configPath = fmt.Sprintf("%v/config.%v.json", NoteClerkData, strings.ToLower(NoteClerkEnv))

// This is the database implementation for the server is injected into a singleton variable. It can be exchanged so
// long as the new database implementation interfaces with the RDBMSAccessor interface.
var db = &DbPostgres{}

// A mock Db implementation
var mockDb = &MockDb{}

// Inject server service into singleton. This can be replaced by any server that implements the NoteClerkServer
// interfaces, which itself implements the NoteServiceServer interface, a gRPC service interface.
var server NoteClerkServer = &Server{}

// Initialize the logger with a set of default settings. Takes a path to a log file, which is set in the config above,
// and opens the file pointed to by the log path.
// RETURNS: error
func InitializeLogger(logPath string) error {
	log.Formatter = &logrus.JSONFormatter{}
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errors.WithMessage(err, ErrMapStr[InitializeLoggerFailsOpenLogFile])
	}

	log.SetLevel(logrus.InfoLevel)
	var writer io.Writer = os.Stdout
	if Environment != "production" {
		writer = io.MultiWriter(os.Stdout, logFile)
		logrus.SetLevel(logrus.DebugLevel)
	}

	log.Out = writer

	return nil
}
