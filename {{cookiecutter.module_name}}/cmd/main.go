package main

import (
	"github.com/Ryanair/gofrlib/frotel"
	"{{cookiecutter.project_name}}/{{cookiecutter.module_name}}/cmd/handler"
	"{{cookiecutter.project_name}}/{{cookiecutter.module_name}}/cmd/internal"
)

var (
	provider     = internal.NewProvider()
	loggerConfig = provider.ProvideLoggerConfig()
)

func main() {
	handler := handler.New(loggerConfig)
	frotel.Start(handler.Handle)
}
