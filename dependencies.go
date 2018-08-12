package main

import "log"

// Pure dependency injection.
type Dependencies struct {
	DB  DbAccessor
	Log log.Logger
}

// Pure dependency injection vector.
var pdi = Dependencies {
	DB:  &DbPostgres{},
	Log: log.Logger{},
}

