package main

type Dependencies struct {
	DB DbAccessor
}

var dependencies = Dependencies {
	DB: &DbPostgres{},
}
